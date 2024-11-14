package service

import (
	"e-commerce/internal/cart/model"
	"e-commerce/internal/cart/repository"
)

type CartService struct {
	repo *repository.CartRepository
}

func NewCartService(repo *repository.CartRepository) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) CreateCart(cart model.Cart) (model.Cart, error) {
	return s.repo.Create(cart)
}

func (s *CartService) GetCart(id int) (model.Cart, error) {
	return s.repo.Get(id)
}
