package services

import (
	"errors"
	"strings"

	"github.com/Rynoo1/spendly/backend/models"
	"github.com/Rynoo1/spendly/backend/repositories"
	"gorm.io/gorm"
)

type CategoriesService struct {
	DB   *gorm.DB
	Repo *repositories.CategoryRepo
}

func NewCategoryService(
	db *gorm.DB,
	repo *repositories.CategoryRepo,
) *CategoriesService {
	return &CategoriesService{
		DB:   db,
		Repo: repo,
	}
}

// CREATE
func (s *CategoriesService) CreateCategory(
	category *models.Category,
) error {
	if strings.TrimSpace(category.Name) == "" {
		return errors.New("category name is required")
	}

	category.Name = strings.TrimSpace(category.Name)

	return s.Repo.Create(category)
}

// READ ONE
func (s *CategoriesService) GetCategory(
	id uint,
) (*models.Category, error) {
	if id == 0 {
		return nil, errors.New("invalid category ID")
	}
	return s.Repo.GetByID(id)
}

// READ ALL
func (s *CategoriesService) GetCategories() ([]models.Category, error) {
	return s.Repo.GetAll()
}

// UPDATE
func (s *CategoriesService) UpdateCategory(
	id uint,
	updates map[string]interface{},
) error {
	if id == 0 {
		return errors.New("invalid category ID")
	}

	if len(updates) == 0 {
		return errors.New("no updates provided")
	}

	for field := range updates {
		if field != "name" {
			return errors.New("invalid category field: " + field)
		}
	}

	if name, ok := updates["name"]; ok {
		value, ok := name.(string)
		if !ok || strings.TrimSpace(value) == "" {
			return errors.New("category name cannot be empty")
		}
		updates["name"] = strings.TrimSpace(value)
	}
	return s.Repo.Update(id, updates)
}

// DELETE
func (s *CategoriesService) DeleteCategory(
	id uint,
) error {
	if id == 0 {
		return errors.New("invalid category ID")
	}
	return s.Repo.Delete(id)
}
