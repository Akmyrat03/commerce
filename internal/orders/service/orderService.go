package service

import (
	"e-commerce/internal/orders/model"
	"e-commerce/internal/orders/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order *model.Order, items []model.OrderItem) error {
	// Begin transaction
	tx, err := s.repo.DB.Begin()
	if err != nil {
		return err
	}

	// Create the order
	if err := s.repo.Create(order); err != nil {
		tx.Rollback()
		return err
	}

	// Add order items
	for _, item := range items {
		item.OrderID = order.ID
		if err := s.repo.AddOrderItem(&item); err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit transaction
	return tx.Commit()
}

func (s *OrderService) GetOrders(userID int) ([]model.Order, error) {
	return s.repo.GetOrders(userID)
}
