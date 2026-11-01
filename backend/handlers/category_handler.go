package handlers

import (
	"strconv"

	"github.com/Rynoo1/spendly/backend/models"
	"github.com/Rynoo1/spendly/backend/services"
	"github.com/gofiber/fiber/v3"
)

// CREATE A NEW CATEGORY
func CreateCategory(c fiber.Ctx, svc *services.CategoriesService) error {
	var category models.Category

	if err := c.Bind().Body(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid category payload",
		})
	}

	if err := svc.CreateCategory(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(category)
}

// READ ONE CATEGORY BY ID
func GetCategory(c fiber.Ctx, svc *services.CategoriesService) error {
	categoryIDStr := c.Params("id")

	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil || categoryID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid category ID",
		})
	}

	category, err := svc.GetCategory(uint(categoryID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(category)
}

// READ ALL CATEGORIES
func GetCategories(c fiber.Ctx, svc *services.CategoriesService) error {
	categories, err := svc.GetCategories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(categories)
}

// UPDATE CATEGORY BY ID
func UpdateCategory(c fiber.Ctx, svc *services.CategoriesService) error {
	categoryIDStr := c.Params("id")

	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid category ID",
		})
	}

	if categoryID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "category ID must be greater than zero",
		})
	}

	var updates map[string]interface{}
	if err := c.Bind().Body(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid category update payload",
		})
	}

	if err := svc.UpdateCategory(uint(categoryID), updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	category, err := svc.GetCategory(uint(categoryID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(category)
}

// DELETE CATEGORY BY ID
func DeleteCategory(c fiber.Ctx, svc *services.CategoriesService) error {
	categoryIDStr := c.Params("id")

	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid category ID",
		})
	}

	if categoryID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "category ID must be greater than zero",
		})
	}

	if err := svc.DeleteCategory(uint(categoryID)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "category deleted successfully",
	})
}
