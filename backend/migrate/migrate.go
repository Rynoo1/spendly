package migrate

import (
	"fmt"

	"github.com/Rynoo1/spendly/backend/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Category{},
		&models.Transaction{},
	)
	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	return nil
}
