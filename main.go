package main

import (
	"url-shortener/router"
)

func Setup() {
	router.Setup()
}

func main() {
	Setup()
	router.Router.Run()
}
