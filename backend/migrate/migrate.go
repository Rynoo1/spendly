package migrate

import (
	"fmt"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate()
	if err != nil {
		return fmt.Errorf("failed to run migrations: %v", err)
	}
	return nil
}
