package main

import (
	"learn-mock/app"
	"learn-mock/database"
	"learn-mock/models"
	"log"
)

func main() {
	db := database.NewDatabase()
	db.AutoMigrate(models.Product{})
	fiberApp := app.Setup(&db)
	log.Fatal(fiberApp.Listen(":3000"))
}
