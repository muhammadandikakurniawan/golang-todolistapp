package main

import (
	"github.com/joho/godotenv"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/container"
	"github.com/muhammadandikakurniawan/golang-todolistapp/src/delivery"
)

func init() {
	godotenv.Load()
}

func main() {
	dependencyContainer := container.InitializeContainer()
	delivery.Run(&dependencyContainer)
}
