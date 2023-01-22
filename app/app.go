package app

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"learn-mock/database"
	"learn-mock/router"
	"net/http"
)

//go:embed swagger
var schemas embed.FS

func Setup(db *database.Database) *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowHeaders: "*",
		AllowOrigins: "*",
	}))
	app.Use("/docs", filesystem.New(filesystem.Config{
		Root:   http.FS(schemas),
		Browse: true,
		Index:  "index.html",
	}))

	api := app.Group("/api")
	{
		router.ProductRouter(api, db)
	}
	return app
}
