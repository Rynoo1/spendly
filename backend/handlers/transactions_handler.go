package handlers

import (
	"strconv"

	"github.com/Rynoo1/spendly/backend/models"
	"github.com/Rynoo1/spendly/backend/services"
	"github.com/gofiber/fiber/v3"
)

// CREATE A NEW TRANSACTION
func CreateTransaction(c fiber.Ctx, svc *services.TransactionService) error {
	var transaction models.Transaction

	if err := c.Bind().Body(&transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid transaction payload",
		})
	}

	if err := svc.CreateTransaction(&transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(transaction)
}

// READ ONE TRANSACTION BY ID
func GetTransactionByID(c fiber.Ctx, svc *services.TransactionService) error {
	transactionIDStr := c.Params("id")

	transactionID, err := strconv.ParseUint(transactionIDStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid transaction ID",
		})
	}

	if transactionID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "transaction ID must be greater than zero",
		})
	}

	transaction, err := svc.GetTransaction(uint(transactionID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(transaction)
}

// READ ALL TRANSACTIONS
func GetTransactions(c fiber.Ctx, svc *services.TransactionService) error {
	transactions, err := svc.GetTransactions()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(transactions)
}

// UPDATE TRANSACTION BY ID
func UpdateTransaction(c fiber.Ctx, svc *services.TransactionService) error {
	transactionIDStr := c.Params("id")

	transactionID, err := strconv.ParseUint(transactionIDStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid transaction ID",
		})
	}

	if transactionID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "transaction ID must be greater than zero",
		})
	}

	var updates map[string]interface{}
	if err := c.Bind().Body(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid transaction update payload",
		})
	}

	if err := svc.UpdateTransaction(uint(transactionID), updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	transaction, err := svc.GetTransaction(uint(transactionID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(transaction)
}

// DElETE TRANSACTION BY ID
func DeleteTransaction(c fiber.Ctx, svc *services.TransactionService) error {
	transactionIDStr := c.Params("id")

	transactionID, err := strconv.ParseUint(transactionIDStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid transaction ID",
		})
	}

	if transactionID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "transaction ID must be greater than zero",
		})
	}

	if err := svc.DeleteTransaction(uint(transactionID)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "transaction deleted successfully",
	})
}

// READ ALL TRANSACTIONS BY CATEGORY ID
func GetTransactionsByCategoryID(c fiber.Ctx, svc *services.TransactionService) error {
	categoryIDStr := c.Params("category_id")

	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid category ID",
		})
	}

	transactions, err := svc.GetTransactionsByCategoryID(uint(categoryID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(transactions)
}
