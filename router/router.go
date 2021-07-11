package router

import (
	"url-shortener/handlers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func SetupRouter() *gin.Engine {

	router := gin.Default()
	router.GET(("/:id"), handlers.GetLongUrl)
	router.POST("/api/url-shortener/v1/url", handlers.GenerateShortUrl)
	return router
}

func Setup() {
	Router = SetupRouter()
}
