package main

import (
	"example.com/cronjobs"
	"example.com/env"
	"example.com/middleware"
	"example.com/mongoose"
	"example.com/postgres"
	"example.com/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	env.LoadEnv()
	mongoose.InitMongoose()
	postgres.InitPostgres()
}

func main() {
	server := gin.Default()
	server.Use(middleware.LogRequests)
	routes.RegisterRoutes(server)
	cronjobs.GetArticles()
	server.Run(":8080")
}