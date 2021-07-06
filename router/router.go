package router

import (
	"go-projects/handlers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func SetupRouter() *gin.Engine {

	// router := gin.New()
	router := gin.Default()
	router.GET("/hello", handlers.UrlHandler)
	// router.GET("/api/url-shortener/v1/url", handlers.CreateShortUrl)
	router.POST("/api/url-shortener/v1/url", handlers.CreateShortUrl)
	return router
}

func Setup() {
	Router = SetupRouter()
}
