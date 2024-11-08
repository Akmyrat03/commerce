package service

import (
	"e-commerce/internal/products/model"
	"e-commerce/internal/products/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) AddProduct(product model.Product) error {
	return s.repo.CreateProduct(product)
}
