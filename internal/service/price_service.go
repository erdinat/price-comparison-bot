package service

import (
	"github.com/erdinat/internProjectGolang/internal/model"
	"github.com/erdinat/internProjectGolang/internal/repository"
)

type ProductPriceDiffService struct {
	Repo *repository.ProductPriceDiffRepository
}

func (s *ProductPriceDiffService) CreateProductPriceDiff(diff *model.ProductPriceDiff) (*model.ProductPriceDiff, error) {
	return s.Repo.CreateProductPriceDiff(diff)
}

func (s *ProductPriceDiffService) GetPriceDiffByID(id int) (*model.ProductPriceDiff, error) {
	return s.Repo.GetPriceDiffByID(id)
}

func (s *ProductPriceDiffService) UpdatePriceDiff(diff *model.ProductPriceDiff) error {
	return s.Repo.UpdatePriceDiff(diff)
}

func (s *ProductPriceDiffService) DeletePriceDiff(id int) error {
	return s.Repo.DeletePriceDiff(id)
}
