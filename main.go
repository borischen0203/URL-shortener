package main

import (
	"url-shortener/config"
	"url-shortener/database"
	"url-shortener/logger"
	"url-shortener/router"

	_ "github.com/joho/godotenv/autoload"
)

func Setup() {
	logger.Setup()
	config.Setup()
	database.Setup()
	router.Setup()
}

func main() {
	Setup()
	router.Router.Run(":" + config.Env.PORT)
}
