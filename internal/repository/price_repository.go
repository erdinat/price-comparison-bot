package repository

import (
	"database/sql"
	"github.com/erdinat/internProjectGolang/internal/model"
)

type ProductPriceDiffRepository struct {
	DB *sql.DB
}

func NewProductPriceDiffRepository(db *sql.DB) *ProductPriceDiffRepository {
	return &ProductPriceDiffRepository{DB: db}
}

func (r *ProductPriceDiffRepository) CreateProductPriceDiff(diff *model.ProductPriceDiff) (*model.ProductPriceDiff, error) {
	result, err := r.DB.Exec("INSERT INTO product_price_diff (product_id, old_price, new_price, created_at, site_id) VALUES (?, ?, ?, ?, ?)",
		diff.ProductID, diff.OldPrice, diff.NewPrice, diff.CreatedAt, diff.SiteID)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	diff.ID = int(id)
	return diff, nil
}

func (r *ProductPriceDiffRepository) GetPreviousPrice(productID int) (float64, error) {
	query := `
		SELECT new_price
		FROM product_price_diff
		WHERE product_id = ?
		ORDER BY created_at DESC
		LIMIT 1`
	row := r.DB.QueryRow(query, productID)

	var oldPrice float64
	err := row.Scan(&oldPrice)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	return oldPrice, nil
}

func (r *ProductPriceDiffRepository) GetPriceDiffByID(id int) (*model.ProductPriceDiff, error) {
	row := r.DB.QueryRow("SELECT * FROM product_price_diff WHERE id = ?", id)
	var diff model.ProductPriceDiff
	err := row.Scan(&diff.ID, &diff.ProductID, &diff.OldPrice, &diff.NewPrice, &diff.CreatedAt, &diff.SiteID)
	if err != nil {
		return nil, err
	}
	return &diff, nil
}

func (r *ProductPriceDiffRepository) UpdatePriceDiff(diff *model.ProductPriceDiff) error {
	_, err := r.DB.Exec("UPDATE product_price_diff SET old_price = ?, new_price = ?, created_at = ?, site_id = ? WHERE id = ?",
		diff.OldPrice, diff.NewPrice, diff.CreatedAt, diff.SiteID, diff.ID)
	return err
}

func (r *ProductPriceDiffRepository) DeletePriceDiff(id int) error {
	_, err := r.DB.Exec("DELETE FROM product_price_diff WHERE id = ?", id)
	return err
}

func (r *ProductPriceDiffRepository) GetAllPriceDiffs() ([]model.ProductPriceDiff, error) {
	rows, err := r.DB.Query("SELECT * FROM product_price_diff")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var diffs []model.ProductPriceDiff
	for rows.Next() {
		var diff model.ProductPriceDiff
		if err := rows.Scan(&diff.ID, &diff.ProductID, &diff.OldPrice, &diff.NewPrice, &diff.CreatedAt, &diff.SiteID); err != nil {
			return nil, err
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}
