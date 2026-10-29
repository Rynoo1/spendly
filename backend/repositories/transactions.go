package repositories

import (
	"github.com/Rynoo1/spendly/backend/models"
	"gorm.io/gorm"
)

type TransactionRepo struct {
	DB *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) *TransactionRepo {
	return &TransactionRepo{
		DB: db,
	}
}

func (r *TransactionRepo) WithTx(tx *gorm.DB) *TransactionRepo {
	return &TransactionRepo{
		DB: tx,
	}
}

// CREATE
func (r *TransactionRepo) Create(
	transaction *models.Transaction,
) error {
	return r.DB.Create(transaction).Error
}

// READ ONE
func (r *TransactionRepo) GetByID(
	id uint,
) (*models.Transaction, error) {
	var transaction models.Transaction

	err := r.DB.First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// READ ALL
func (r *TransactionRepo) GetAll() ([]models.Transaction, error) {
	var transactions []models.Transaction

	err := r.DB.Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// UPDATE
func (r *TransactionRepo) Update(
	id uint,
	updates map[string]interface{},
) error {
	result := r.DB.Model(&models.Transaction{}).Where("id = ?", id).Updates(updates)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// DELETE
func (r *TransactionRepo) Delete(id uint) error {
	result := r.DB.Delete(&models.Transaction{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
