package services

import (
	"github.com/Rynoo1/spendly/backend/repositories"
	"gorm.io/gorm"
)

type AppServices struct {
	DB *gorm.DB

	CategoryRepo    *repositories.CategoryRepo
	TransactionRepo *repositories.TransactionRepo

	CategoryService    *CategoriesService
	TransactionService *TransactionService
}

func NewAppServices(db *gorm.DB) *AppServices {
	categoryRepo := repositories.NewCategoryRepo(db)
	transactionRepo := repositories.NewTransactionRepo(db)

	return &AppServices{
		DB: db,

		CategoryRepo:    categoryRepo,
		TransactionRepo: transactionRepo,

		CategoryService:    NewCategoryService(db, categoryRepo),
		TransactionService: NewTransactionService(db, transactionRepo),
	}
}
