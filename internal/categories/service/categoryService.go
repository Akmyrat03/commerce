package service

import (
	"e-commerce/internal/categories/model"
	"e-commerce/internal/categories/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Create(category *model.Category) (int, error) {
	return s.repo.Create(category)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *CategoryService) Update(id int, name string) error {
	return s.repo.Update(id, name)
}
