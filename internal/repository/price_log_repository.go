package repository

import (
	"database/sql"
	"github.com/erdinat/internProjectGolang/internal/model"
)

type ProductPriceDiffLogRepository struct {
	DB *sql.DB
}

func NewProductPriceDiffLogRepository(db *sql.DB) *ProductPriceDiffLogRepository {
	return &ProductPriceDiffLogRepository{}
}

func (r *ProductPriceDiffLogRepository) CreateProductPriceDiffLog(log *model.ProductPriceDiffLog) (*model.ProductPriceDiffLog, error) {
	result, err := r.DB.Exec("INSERT INTO product_price_diff_log (product_id, data, created_at) VALUES (?, ?, ?)",
		log.ProductID, log.Data, log.CreatedAt)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	log.ID = int(id)
	return log, nil
}

func (r *ProductPriceDiffLogRepository) GetProductPriceDiffLogByID(id int) (*model.ProductPriceDiffLog, error) {
	row := r.DB.QueryRow("SELECT * FROM product_price_diff_log WHERE id = ?", id)
	var log model.ProductPriceDiffLog
	err := row.Scan(&log.ID, &log.ProductID, &log.Data, &log.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &log, nil
}

func (r *ProductPriceDiffLogRepository) UpdateProductPriceDiffLog(log *model.ProductPriceDiffLog) error {
	_, err := r.DB.Exec("UPDATE product_price_diff_log SET data = ?, created_at = ? WHERE id = ?",
		log.Data, log.CreatedAt, log.ID)
	return err
}

func (r *ProductPriceDiffLogRepository) DeleteProductPriceDiffLog(id int) error {
	_, err := r.DB.Exec("DELETE FROM product_price_diff_log WHERE id = ?", id)
	return err
}

func (r *ProductPriceDiffLogRepository) GetAllProductPriceDiffLogs() ([]model.ProductPriceDiffLog, error) {
	rows, err := r.DB.Query("SELECT * FROM product_price_diff_log")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []model.ProductPriceDiffLog
	for rows.Next() {
		var log model.ProductPriceDiffLog
		if err := rows.Scan(&log.ID, &log.ProductID, &log.Data, &log.CreatedAt); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}
