package repository

import (
	"e-commerce/internal/products/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	Products = "products"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(product model.Product) error {
	query := fmt.Sprintf("INSERT INTO %s (name, description, image, price, category_id) VALUES ($1, $2, $3, $4, $5)", Products)
	_, err := r.db.Exec(query, product.Name, product.Description, product.Image, product.Price, product.CategoryID)
	if err != nil {
		return err
	}

	return nil
}
