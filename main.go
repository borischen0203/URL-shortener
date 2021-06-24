package main

import (
	"go-projects/router"
)

func Setup() {
	router.Setup()
}

func main() {
	Setup()
	router.Router.Run()
}
