package repository

import (
	"e-commerce/internal/products/model"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	Products      = "products"
	LikedProducts = "liked_products"
	ProductImages = "product_images"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// Admin can create a new product
func (r *ProductRepository) CreateProduct(product model.Product) (int, error) {

	tx, err := r.db.Begin()
	if err != nil {
		log.Println("Error starting transaction:", err)
		return 0, err
	}

	var productID int
	query := fmt.Sprintf("INSERT INTO %s (name, description, price, status, category_id) VALUES ($1, $2, $3, $4, $5) RETURNING id", Products)
	log.Println("Executing product query:", query)
	log.Println("Values:", product.Name, product.Description, product.Price, product.Status, product.CategoryID)

	err = tx.QueryRow(query, product.Name, product.Description, product.Price, product.Status, product.CategoryID).Scan(&productID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	log.Println("Generated Product ID:", productID)

	imageQuery := fmt.Sprintf("INSERT INTO %v (product_id, image_url) VALUES ($1, $2)", ProductImages)
	for _, image := range product.Images {
		log.Println("Inserting image:", image)
		_, err := tx.Exec(imageQuery, product.ID, image)
		if err != nil {
			log.Println("Error inserting image:", err)
			tx.Rollback()
			return 0, err
		}
	}

	if err := tx.Commit(); err != nil {
		log.Println("Error committing transaction:", err)
		return 0, err
	}

	log.Println("Transaction committed successfully.")

	return productID, nil
}

func (r *ProductRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %v WHERE id = $1", Products)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

// Admin can see all drafted and published products
func (r *ProductRepository) GetAll() ([]model.Product, error) {
	var products []model.Product
	query := `SELECT p.id, p.name, c.name AS category_name, p.description, p.price, p.image, p.status, p.created_at FROM products AS p
	INNER JOIN categories AS c ON c.id= p.category_id
	ORDER BY p.id ASC`
	err := r.db.Select(&products, query)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// User can see all published products
func (r *ProductRepository) GetAllPublishedProducts() ([]model.Product, error) {
	var products []model.Product
	query := `SELECT p.id, p.name, c.name AS category_name, p.description, p.price, p.image, p.status, p.created_at FROM products AS p
	INNER JOIN categories AS c ON c.id= p.category_id
	WHERE p.status=$1
	ORDER BY p.id ASC`
	err := r.db.Select(&products, query, "published")
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProductByCategory(categoryName string) ([]model.Product, error) {
	var products []model.Product
	query := `SELECT p.id, p.name, p.description, p.price, p.image, p.status, p.created_at, c.name AS category_name FROM products AS p
	INNER JOIN categories AS c ON c.id=p.category_id
	WHERE c.name=$1`
	err := r.db.Select(&products, query, categoryName)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(id int) (model.Product, error) {
	var product model.Product
	query := `SELECT p.name, p.description, p.price, p.image, p.status, p.created_at, c.name AS category_name FROM products AS p
	INNER JOIN categories AS c ON c.id=p.category_id
	WHERE p.id=$1`
	err := r.db.Get(&product, query, id)
	if err != nil {
		return model.Product{}, nil
	}

	return product, nil
}

func (r *ProductRepository) LikeProduct(userID, productID int) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, product_id) VALUES ($1, $2) ON CONFLICT DO NOTHING", LikedProducts)
	_, err := r.db.Exec(query, userID, productID)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) UnlikeProduct(userID, productID int) error {
	query := fmt.Sprintf("DELETE FROM %v WHERE user_id = $1 AND product_id = $2", LikedProducts)
	_, err := r.db.Exec(query, userID, productID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) GetLikedProducts(userID int) ([]model.LikedProduct, error) {
	var likedProducts []model.LikedProduct
	query := fmt.Sprintf("SELECT id, user_id, product_id FROM %s WHERE user_id = $1", LikedProducts)
	err := r.db.Select(&likedProducts, query, userID)
	if err != nil {
		return nil, err
	}

	return likedProducts, nil
}
