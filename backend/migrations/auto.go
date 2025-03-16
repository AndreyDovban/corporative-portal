package main

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate()
}
