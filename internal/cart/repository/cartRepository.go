package repository

import (
	"e-commerce/internal/cart/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	Cart = "shopping_cart"
)

type CartRepository struct {
	db *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) Create(cart model.Cart) (model.Cart, error) {
	query := fmt.Sprintf("INSERT INTO %s (user_id, created_at) VALUES ($1, Now()) RETURNING id", Cart)
	err := r.db.QueryRow(query, cart.UserID).Scan(&cart.ID)
	if err != nil {
		return model.Cart{}, err
	}

	return cart, nil
}

func (r *CartRepository) Get(id int) (model.Cart, error) {
	var cart model.Cart
	query := fmt.Sprintf("SELECT id, user_id, created_at FROM %s WHERE id = $1", Cart)
	err := r.db.QueryRow(query, id).Scan(&cart.ID, &cart.UserID, &cart.CreatedAt)
	if err != nil {
		return model.Cart{}, err
	}

	return cart, nil
}
