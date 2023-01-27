package router

import (
	"github.com/gofiber/fiber/v2"
	"learn-mock/database"
	"learn-mock/handlers"
	"learn-mock/repository"
)

func ProductRouter(router fiber.Router, db *database.Database) {
	group := router.Group("/product")
	repo := repository.NewProductRepositoryImpl(db)
	handler := handlers.NewProductHandler(repo)
	{
		group.Post("/", handler.Create())
		group.Get("/", handler.Get())
		group.Delete("/:id", handler.Delete())
		group.Patch("/:id", handler.Update())
	}
}
