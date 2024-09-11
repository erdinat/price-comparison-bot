package scraper

import (
	"database/sql"
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
)

type ScrapedProduct struct {
	ProductName string
	SKU         string
	Price       string
	Image       string
	SiteName    string
}

func ScrapeSite(siteName, url string, db *sql.DB) []ScrapedProduct {
	var products []ScrapedProduct

	c := colly.NewCollector()

	switch siteName {
	case "Amazon":
		c.OnHTML(".s-product-container", func(e *colly.HTMLElement) {
			sku := e.Attr("data-product-sku") //
			title := e.ChildText(".product-title")
			price := e.ChildText(".price")
			image := e.ChildAttr(".image img", "src")
			if sku != "" && title != "" && price != "" && image != "" {
				product := ScrapedProduct{
					ProductName: title,
					SKU:         sku,
					Price:       price,
					Image:       image,
					SiteName:    siteName,
				}
				products = append(products, product)
			}
		})
	case "Özdilek":
		c.OnHTML(".ozdilek-product-container", func(e *colly.HTMLElement) {
			title := e.ChildText(".product-title")
			price := e.ChildText(".price")
			sku := e.Attr("data-sku")
			image := e.ChildAttr(".product-img img", "src")
			if title != "" && price != "" && sku != "" && image != "" {
				product := ScrapedProduct{
					ProductName: title,
					SKU:         sku,
					Price:       price,
					Image:       image,
					SiteName:    siteName,
				}
				products = append(products, product)
			}
		})
	case "Pazarama":
		c.OnHTML(".pazarama-product-container", func(e *colly.HTMLElement) {
			title := e.ChildText(".product-name")
			price := e.ChildText(".product-price")
			sku := e.Attr("data-sku")
			image := e.ChildAttr(".product-image img", "src")
			if title != "" && price != "" && sku != "" && image != "" {
				product := ScrapedProduct{
					ProductName: title,
					SKU:         sku,
					Price:       price,
					Image:       image,
					SiteName:    siteName,
				}
				products = append(products, product)
			}
		})
	default:
		fmt.Println("Bu site için scraping stratejisi tanımlanmadı:", siteName)
	}

	err := c.Visit(url)
	if err != nil {
		log.Fatalf("Site ziyaret edilirken hata oluştu: %v", err)
	}

	return products
}
