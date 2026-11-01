package routes

import (
	"github.com/Rynoo1/spendly/backend/handlers"
	"github.com/Rynoo1/spendly/backend/repositories"
	"github.com/Rynoo1/spendly/backend/services"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, svc *services.AppServices, db *gorm.DB) {
	if svc == nil {
		svc = &services.AppServices{}
	}

	if svc.CategoryRepo == nil {
		svc.CategoryRepo = repositories.NewCategoryRepo(db)
	}
	if svc.TransactionRepo == nil {
		svc.TransactionRepo = repositories.NewTransactionRepo(db)
	}

	categoryService := services.NewCategoryService(db, svc.CategoryRepo)
	transactionService := services.NewTransactionService(db, svc.TransactionRepo)

	api := app.Group("/api")

	categories := api.Group("/categories")
	categories.Get("/", func(c fiber.Ctx) error {
		return handlers.GetCategories(c, categoryService)
	})
	categories.Post("/", func(c fiber.Ctx) error {
		return handlers.CreateCategory(c, categoryService)
	})
	categories.Get("/:id", func(c fiber.Ctx) error {
		return handlers.GetCategory(c, categoryService)
	})
	categories.Put("/:id", func(c fiber.Ctx) error {
		return handlers.UpdateCategory(c, categoryService)
	})
	categories.Delete("/:id", func(c fiber.Ctx) error {
		return handlers.DeleteCategory(c, categoryService)
	})

	transactions := api.Group("/transactions")
	transactions.Post("/", func(c fiber.Ctx) error {
		return handlers.CreateTransaction(c, transactionService)
	})
	transactions.Get("/", func(c fiber.Ctx) error {
		return handlers.GetTransactions(c, transactionService)
	})
	transactions.Get("/:id", func(c fiber.Ctx) error {
		return handlers.GetTransactionByID(c, transactionService)
	})
	transactions.Put("/:id", func(c fiber.Ctx) error {
		return handlers.UpdateTransaction(c, transactionService)
	})
	transactions.Delete("/:id", func(c fiber.Ctx) error {
		return handlers.DeleteTransaction(c, transactionService)
	})
	transactions.Get("/category/:category_id", func(c fiber.Ctx) error {
		return handlers.GetTransactionsByCategoryID(c, transactionService)
	})

	app.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})
}
