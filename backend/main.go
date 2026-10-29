package main

import (
	"fmt"
	"log"

	"github.com/Rynoo1/spendly/backend/config"
	"github.com/Rynoo1/spendly/backend/migrate"
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
}
