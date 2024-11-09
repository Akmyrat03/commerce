package repository

import (
	"e-commerce/internal/categories/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	Category = "categories"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(category *model.Category) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %v (name) VALUES ($1) RETURNING id", Category)
	row := r.db.QueryRow(query, category.Name)
	if err := row.Scan(&id); err != nil {
		fmt.Println("Error during row.Scan: ", err)
		return 0, err
	}

	return id, nil
}

func (r *CategoryRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %v WHERE id = $1", Category)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *CategoryRepository) Update(id int, name string) error {
	query := fmt.Sprintf("UPDATE %v SET name = $1 WHERE id = $2", Category)
	_, err := r.db.Exec(query, name, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *CategoryRepository) GetAll() ([]model.Category, error) {
	var categories []model.Category
	query := fmt.Sprintf("SELECT id, name, created_at FROM %v", Category)
	err := r.db.Select(&categories, query)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
