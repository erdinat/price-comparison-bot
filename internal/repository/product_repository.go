package repository

import (
	"database/sql"
	"github.com/erdinat/internProjectGolang/internal/model"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}

}

func (r *ProductRepository) GetProductByID(id int) (*model.Product, error) {
	row := r.DB.QueryRow("SELECT * FROM products WHERE id = ?", id)
	var p model.Product
	err := row.Scan(&p.ID, &p.ProductName, &p.SKU, &p.Barcode, &p.CreatedAt, &p.Status, &p.ProductImage)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) CreateProduct(p *model.Product) (*model.Product, error) {
	_, err := r.DB.Exec("INSERT INTO products (product_name,sku,barcode,created_at,status, product_image) VALUES (?, ?,?,?,?,?)",
		p.ProductName, p.SKU, p.Barcode, p.CreatedAt, p.Status, p.ProductImage)
	return p, err
}

func (r *ProductRepository) UpdateProduct(p model.Product) (*model.Product, error) {
	_, err := r.DB.Exec("Update products set product_name=?,sku=?,barcode = ?, created_at = ?, status = ?,product_image=? where id=?",
		p.ProductName, p.SKU, p.Barcode, p.CreatedAt, p.Status, p.ProductImage, p.ID)
	return &p, err
}

func (r *ProductRepository) DeleteProduct(id int) error {
	_, err := r.DB.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}

func (r *ProductRepository) GetAllProducts() ([]model.Product, error) {
	rows, err := r.DB.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.ProductName, &p.SKU, &p.Barcode, &p.CreatedAt, &p.Status, &p.ProductImage); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
