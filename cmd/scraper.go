package cmd

import (
	"database/sql"
	"fmt"
	"github.com/erdinat/internProjectGolang/internal/model"
	"github.com/erdinat/internProjectGolang/internal/repository"
	"github.com/gocolly/colly/v2"
	"log"
	"strconv"
	"time"
)

func ScraperSites(db *sql.DB) {
	siteRepo := repository.NewSiteRepository(db)
	sites, err := siteRepo.GetAllSites()
	if err != nil {
		log.Fatal("Sites tablosundan veri alırken hata oluştu: %v", err)
	}

	for _, site := range sites {
		fmt.Println("Scrapping site: %s\n", site.URL)

		c := colly.NewCollector()

		c.OnHTML(".js-product-vertical", func(e *colly.HTMLElement) {
			// Ürün bilgilerini çek
			sku := e.Attr("data-product-sku")
			title := e.ChildText(".product__name")
			price := e.ChildText(".product__prices-sale")
			image := e.ChildAttr(".product__image img", "src")
			barcode := e.ChildText(".product__barcode")

			if sku == "" || title == "" || price == "" || image == "" || barcode == "" {
				fmt.Println("Ürün atlandı.")
				return
			}

			priceFloat, err := strconv.ParseFloat(price, 64)
			if err != nil {
				log.Printf("Fiyat değeri dönüştürülürken hata oluştu: %v", err)
				return
			}

			productRepo := repository.NewProductRepository(db)
			product := &model.Product{
				ProductName:  title,
				SKU:          sku,
				Barcode:      barcode,
				CreatedAt:    time.Now(),
				Status:       1,
				ProductImage: image,
			}

			createdProduct, err := productRepo.CreateProduct(product)
			if err != nil {
				log.Fatal("Ürün eklenirken hata oluştu: %v", err)
				return
			} else {
				fmt.Println("Ürün eklendi: %s\n", createdProduct)
			}

			priceDiffRepo := repository.NewProductPriceDiffRepository(db)
			oldPrice, err := priceDiffRepo.GetPreviousPrice(createdProduct.ID)
			priceDiff := &model.ProductPriceDiff{
				ProductID: createdProduct.ID,
				OldPrice:  oldPrice,
				NewPrice:  priceFloat,
				CreatedAt: time.Now(),
				SiteID:    site.ID,
			}

			createdPriceDiff, err := priceDiffRepo.CreateProductPriceDiff(priceDiff)
			if err != nil {
				log.Fatal("Price Diff eklenirken hata oluştu: %v", err)
				return
			} else {
				fmt.Printf("Price diff eklendi: %s için %s\n", createdPriceDiff, price)
			}

			priceDiffLogRepo := repository.NewProductPriceDiffLogRepository(db)
			priceDiffLog := &model.ProductPriceDiffLog{
				ProductID: product.ID,
				Data:      fmt.Sprintf("Ürün %s için fiyat değişti", title),
				CreatedAt: time.Now(),
			}

			_, err = priceDiffLogRepo.CreateProductPriceDiffLog(priceDiffLog)
			if err != nil {
				log.Printf("Price diff log eklenirken hata oluştu: %v", err)
				return
			}
			fmt.Printf("Price diff log eklendi: %s için\n", title)

		})

		err = c.Visit(site.URL)
		if err != nil {
			log.Fatal("Site ziyaret edilirken hata oluştu: %w", err)
		}
	}
}
