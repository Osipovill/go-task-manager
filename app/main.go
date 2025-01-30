// /backend/app/main.go
package main

import (
	"app/api/routes"
	"app/config"
	"app/db"
	"app/dependencies"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация конфигурации
	config.Init()

	// Инициализация базы данных
	db.Init()

	// Создание маршрутизатора Gin
	r := gin.Default()

	// Инициализация зависимостей
	dependencies.Init()

	// Настройка роутеров
	routes.SetupRoutes(r)

	// Запуск приложения
	r.Run(":8080")
}
