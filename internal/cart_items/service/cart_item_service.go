package service

import (
	"e-commerce/internal/cart_items/model"
	"e-commerce/internal/cart_items/repository"
)

type CartItemService struct {
	repo *repository.CartItemRepository
}

func NewCartItemService(repo *repository.CartItemRepository) *CartItemService {
	return &CartItemService{repo: repo}
}

func (service *CartItemService) Create(item model.CartItem) (model.CartItem, error) {
	return service.repo.Create(item)
}
