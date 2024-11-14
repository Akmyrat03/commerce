package repository

import (
	"e-commerce/internal/cart_items/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	cartItem = "cart_items"
)

type CartItemRepository struct {
	db *sqlx.DB
}

func NewCartItemRepository(db *sqlx.DB) *CartItemRepository {
	return &CartItemRepository{db: db}
}

func (r *CartItemRepository) Create(item model.CartItem) (model.CartItem, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (cart_id, product_id, quantity) VALUES ($1, $2, $3) RETURNING id", cartItem)
	err := r.db.QueryRow(query, item.CartID, item.ProductID, item.Quantity).Scan(&id)
	if err != nil {
		return model.CartItem{}, err
	}

	return item, nil

}
