package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	routes := server.Group("api")
	routes.GET("/articles", getArticles)
	routes.POST("/article", createArticle)
}
