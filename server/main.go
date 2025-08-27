package main

import (
		"github.com/gofiber/fiber/v2"
		"gorm.io/driver/mysql"
		"gorm.io/gorm"
		"os"
		"github.com/tupap1/listed/server/models"
	)

func main() {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		db.AutoMigrate(&models.User{})
		return c.SendString("tablas creadas")
	})

	app.Listen(":8080")
}