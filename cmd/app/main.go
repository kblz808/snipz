package main

import (
	"context"
	"log"
	"snipz/internal/storage"
	"snipz/internal/utils"
)

func main() {
	config, err := utils.New()
	if err != nil {
		log.Fatal("error loading env variables", err)
	}

	ctx := context.Background()
	db, err := storage.New(ctx, config.DB)
	if err != nil {
		log.Fatal("error initializing database connection: ", err)
	}
	defer db.Close()

	log.Println("successfully connected to the database")

	err = db.Migrate()
	if err != nil {
		log.Fatal("error migrating database", err)
	}

	log.Println("successfully migrated the database")
}
