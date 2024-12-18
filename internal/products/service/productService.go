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

func (s *ProductService) AddProduct(product model.Product) (int, error) {
	return s.repo.CreateProduct(product)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.repo.Delete(id)
}

func (s *ProductService) GetAll() ([]model.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) GetAllPublishedProducts() ([]model.Product, error) {
	return s.repo.GetAllPublishedProducts()
}

func (s *ProductService) GetProductByCatName(category string) ([]model.Product, error) {
	return s.repo.GetProductByCategory(category)
}

func (s *ProductService) GetProductByID(id int) (model.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s *ProductService) LikeProduct(userID, productID int) error {
	return s.repo.LikeProduct(userID, productID)
}

func (s *ProductService) UnlikeProduct(userID, productID int) error {
	return s.repo.UnlikeProduct(userID, productID)
}

func (s *ProductService) GetLikedProducts(userID int) ([]model.LikedProduct, error) {
	return s.repo.GetLikedProducts(userID)
}
