package service

import (
	"github.com/erdinat/internProjectGolang/internal/model"
	"github.com/erdinat/internProjectGolang/internal/repository"
)

type ProductService struct {
	Repo *repository.ProductRepository
}

func (s *ProductService) CreateProduct(product *model.Product) (*model.Product, error) {
	return s.Repo.CreateProduct(product)
}

func (s *ProductService) GetProductByID(id int) (*model.Product, error) {
	return s.Repo.GetProductByID(id)
}

func (s *ProductService) UpdateProduct(product *model.Product) (*model.Product, error) {
	return s.Repo.UpdateProduct(*product)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.Repo.DeleteProduct(id)
}

func (s *ProductService) GetAllProducts() ([]model.Product, error) {
	return s.Repo.GetAllProducts()
}

func (s *ProductService) GetAllPriceDiff() (interface{}, interface{}) {
	return s.Repo.GetAllProducts()
}
