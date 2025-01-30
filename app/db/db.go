package db

import (
	"app/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	// Открытие соединения с SQLite через GORM
	DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	// Миграция: создание таблицы Task, если она еще не существует
	err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}
	log.Println("Database Initialized")
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("failed to get sql.DB from gorm: ", err)
	}
	sqlDB.Close()
}
