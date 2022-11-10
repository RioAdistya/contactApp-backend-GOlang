package repository

import "github.com/gofiber/fiber/v2"

func (repo *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/users", repo.GetUsers)
	api.Post("/users", repo.CreateUsers)
	api.Patch("/users/:id", repo.UpdateUsers)
	api.Delete("/users/:id", repo.DeleteUsers)
	api.Get("/users/:id", repo.GetUserByID)
}
