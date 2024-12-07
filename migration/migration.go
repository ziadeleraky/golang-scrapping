package main

import (
	"log"

	"example.com/env"
	"example.com/models"
	"example.com/mongoose"
	"example.com/postgres"
)

func init() {
	env.LoadEnv()
	mongoose.InitMongoose()
	postgres.InitPostgres()
}

func main() {
	err := postgres.PostgresDb.AutoMigrate(&models.Article{})
	if err != nil {
		log.Fatal("Error migrating to database", err)
	}
	log.Print("Migrated Succesfully")
}
