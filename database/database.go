package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"learn-mock/models"
	"os"
)

type Database struct {
	*gorm.DB
}

func NewDatabase() Database {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_DB")
	pass := os.Getenv("POSTGRES_USER")
	dbname := os.Getenv("POSTGRES_PASSWORD")
	dsn := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, pass, dbname)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		panic(err)
	}
	return Database{db}
}
