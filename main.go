package main

import (
	"learn-mock/app"
	"learn-mock/database"
	"log"
)

func main() {
	db := database.NewDatabase()
	fiberApp := app.Setup(&db)
	log.Fatal(fiberApp.Listen(":3000"))
}
