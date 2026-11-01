package main

import (
	"fmt"
	"log"

	"github.com/Rynoo1/spendly/backend/config"
	"github.com/Rynoo1/spendly/backend/migrate"
	"github.com/Rynoo1/spendly/backend/routes"
	"github.com/Rynoo1/spendly/backend/services"
	"github.com/gofiber/fiber/v3"
)

func main() {

	db, err := config.InitDb()
	if err != nil {
		log.Fatalf("Database init failed: %v", err)
	}

	sqlDB, _ := db.DB()
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}
	log.Println("Database connection OK")

	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Printf("error closing database: %v", err)
		}
	}()

	if err := migrate.RunMigrations(db); err != nil {
		log.Fatalf("Migration failed: %v:", err)
	}
	fmt.Println("Database migrated successfully")

	app := fiber.New()
	appServices := services.NewAppServices(db)

	routes.SetupRoutes(app, appServices, db)

	log.Println("Starting spendly API server on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
