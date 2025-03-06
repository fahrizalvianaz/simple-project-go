package main

import (
	"bookstore-framework/configs"
	"bookstore-framework/migrations"
	"bookstore-framework/pkg"
	"bookstore-framework/routes"
	"log"
)

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config : %v", err)
	}

	db, err := pkg.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := migrations.Migrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	router := routes.Router(db)
	router.Run(":8080")
}
