package main

import (
	"github.com/RioAdistya/contactApp-go-react/bootstrap"
	"github.com/RioAdistya/contactApp-go-react/repository"
	"github.com/gofiber/fiber/v2"
)

type Repository repository.Repository

func main() {
	app := fiber.New()
	bootstrap.InitializeApp(app)
}
