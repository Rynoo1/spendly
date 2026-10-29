package services

import (
	"errors"
	"strings"

	"github.com/Rynoo1/spendly/backend/models"
	"github.com/Rynoo1/spendly/backend/repositories"
	"gorm.io/gorm"
)

type TransactionService struct {
	DB   *gorm.DB
	Repo *repositories.TransactionRepo
}

func NewTransactionService(
	db *gorm.DB,
	repo *repositories.TransactionRepo,
) *TransactionService {
	return &TransactionService{
		DB:   db,
		Repo: repo,
	}
}

// CREATE
func (s *TransactionService) CreateTransaction(
	transaction *models.Transaction,
) error {
	if strings.TrimSpace(transaction.Description) == "" {
		return errors.New("description is required")
	}

	if transaction.AmountCents <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if transaction.CategoryID == 0 {
		return errors.New("category is required")
	}

	return s.Repo.Create(transaction)
}

// READ ONE
func (s *TransactionService) GetTransaction(
	id uint,
) (*models.Transaction, error) {
	if id == 0 {
		return nil, errors.New("invalid transaction ID")
	}

	return s.Repo.GetByID(id)
}

// READ ALL
func (s *TransactionService) GetTransactions() ([]models.Transaction, error) {
	return s.Repo.GetAll()
}

// UPDATE
func (s *TransactionService) UpdateTransaction(
	id uint,
	updates map[string]interface{},
) error {
	if id == 0 {
		return errors.New("invalid transaction ID")
	}

	if len(updates) == 0 {
		return errors.New("no updates provided")
	}

	allowedFields := map[string]bool{
		"description":  true,
		"amount_cents": true,
		"category_id":  true,
		"date":         true,
	}

	for field := range updates {
		if !allowedFields[field] {
			return errors.New("invalid update field: " + field)
		}
	}

	return s.Repo.Update(id, updates)
}

// DELETE
func (s *TransactionService) DeleteTransaction(
	id uint,
) error {
	if id == 0 {
		return errors.New("invalid transaction ID")
	}
	return s.Repo.Delete(id)
}
