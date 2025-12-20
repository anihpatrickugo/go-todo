package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"my-todo/models"
)


var DB *gorm.DB

func Connect() {
	var err error

	DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	log.Println("Database connected")
}


func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Todo{},
	)
	if err != nil {
		log.Fatal("migration failed:", err)
	}
}