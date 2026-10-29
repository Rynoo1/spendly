package main

import (
	"log"
)

func main() {

	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Database init failed: %v", err)
	}
}
