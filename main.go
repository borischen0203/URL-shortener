package main

import (
	"url-shortener/config"
	"url-shortener/database"
	"url-shortener/router"
)

func Setup() {
	config.Setup()
	database.Setup()
	router.Setup()
}

func main() {
	Setup()
	router.Router.Run()
}
