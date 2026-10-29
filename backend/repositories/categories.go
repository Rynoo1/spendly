package repositories

import (
	"github.com/Rynoo1/spendly/backend/models"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	DB *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{
		DB: db,
	}
}

func (r *CategoryRepo) WithTx(tx *gorm.DB) *CategoryRepo {
	return &CategoryRepo{
		DB: tx,
	}
}

// CREATE
func (r *CategoryRepo) Create(
	category *models.Category,
) error {
	return r.DB.Create(category).Error
}

// READ ONE
func (r *CategoryRepo) GetByID(
	id uint,
) (*models.Category, error) {
	var category models.Category

	err := r.DB.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// READ ALL
func (r *CategoryRepo) GetAll() ([]models.Category, error) {
	var categories []models.Category

	err := r.DB.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// UPDATE
func (r *CategoryRepo) Update(
	id uint,
	updates map[string]interface{},
) error {
	result := r.DB.Model(&models.Category{}).Where("id = ?", id).Updates(updates)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// DELETE
func (r *CategoryRepo) Delete(id uint) error {
	result := r.DB.Delete(&models.Category{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
