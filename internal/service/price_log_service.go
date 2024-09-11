package service

import (
	"github.com/erdinat/internProjectGolang/internal/model"
	"github.com/erdinat/internProjectGolang/internal/repository"
)

type ProductPriceDiffLogService struct {
	Repo *repository.ProductPriceDiffLogRepository
}

func (s *ProductPriceDiffLogService) CreateProductPriceDiffLog(log *model.ProductPriceDiffLog) (*model.ProductPriceDiffLog, error) {
	return s.Repo.CreateProductPriceDiffLog(log)
}

func (s *ProductPriceDiffLogService) GetProductPriceDiffLogByID(id int) (*model.ProductPriceDiffLog, error) {
	return s.Repo.GetProductPriceDiffLogByID(id)
}

func (s *ProductPriceDiffLogService) UpdateProductPriceDiffLog(log *model.ProductPriceDiffLog) error {
	return s.Repo.UpdateProductPriceDiffLog(log)
}

func (s *ProductPriceDiffLogService) DeleteProductPriceDiffLog(id int) error {
	return s.Repo.DeleteProductPriceDiffLog(id)
}

func (s *ProductPriceDiffLogService) GetAllProductPriceDiffLogs() ([]model.ProductPriceDiffLog, error) {
	return s.Repo.GetAllProductPriceDiffLogs()
}
