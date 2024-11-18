package repository

import (
	"e-commerce/internal/users/model"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	Users = "users"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user *model.User) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (username, phone_number, password) VALUES ($1, $2, $3) RETURNING id`, Users)
	row := r.db.QueryRow(query, user.Username, user.PhoneNumber, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepository) GetUser(username, password string) (model.User, error) {
	query := fmt.Sprintf(`SELECT id, username, phone_number, password, role FROM %v WHERE username= $1 AND password=$2`, Users)

	var user model.User
	err := r.db.Get(&user, query, username, password)
	if err != nil {
		return model.User{}, errors.New("incorrect username or password")
	}

	return user, nil
}

func (r *UserRepository) GetUserByField(field, value string) (model.User, error) {
	if field != "username" && field != "phone_number" {
		return model.User{}, fmt.Errorf("unsupported field: %s", field)
	}

	query := fmt.Sprintf("SELECT id, username, phone_number, password FROM %v WHERE %s= $1", Users, field)
	var user model.User
	err := r.db.Get(&user, query, value)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepository) DeleteUser(userID int) error {
	query := fmt.Sprintf(`DELETE FROM %v WHERE id= $1`, Users)
	_, err := r.db.Exec(query, userID)
	return err
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	query := fmt.Sprintf("SELECT id, username, phone_number, password, role, created_at, updated_at FROM %v", Users)
	var users []model.User
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}
