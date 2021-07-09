package main

import (
	"url-shortener/database"
	"url-shortener/router"
)

func Setup() {
	router.Setup()
	database.Setup()
}

func main() {
	Setup()
	router.Router.Run()
}
