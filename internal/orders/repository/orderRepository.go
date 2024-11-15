package repository

import (
	"e-commerce/internal/orders/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	Orders     = "orders"
	OrderItems = "order_items"
)

type OrderRepository struct {
	DB *sqlx.DB
}

func NewOrderRepository(DB *sqlx.DB) *OrderRepository {
	return &OrderRepository{DB: DB}
}

// Create creates a new order
func (r *OrderRepository) Create(order *model.Order) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, total_price, status) VALUES ($1, $2, $3) RETURNING id, created_at", Orders)
	err := r.DB.QueryRow(query, order.UserID, order.TotalPrice, order.Status).Scan(&order.ID, &order.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// AddOrderItem adds a new order item
func (r *OrderRepository) AddOrderItem(item *model.OrderItem) error {
	query := fmt.Sprintf("INSERT INTO %s (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4) RETURNING id", OrderItems)
	err := r.DB.QueryRow(query, item.OrderID, item.ProductID, item.Quantity, item.Price).Scan(&item.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetOrders retrieves all orders for a user
func (r *OrderRepository) GetOrders(userID int) ([]model.Order, error) {
	query := fmt.Sprintf("SELECT id, user_id, total_price, status, created_at FROM %s WHERE user_id = $1", Orders)
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []model.Order{}

	for rows.Next() {
		var order model.Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.TotalPrice, &order.Status, &order.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil

}
