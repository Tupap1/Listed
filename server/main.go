package main

import (
		"fmt"
		"github.com/gofiber/fiber/v2"
		"gorm.io/driver/mysql"
		"gorm.io/gorm"
		//"os"
		"github.com/Tupap1/Listed/server/models"
	)

func main() {
	dsn := os.Getenv("root:7iu7Wi0@tcp(127.0.0.1:3306)/adm?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect databaseeee:", err)
	}


	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		db.AutoMigrate(&models.User{})
		return c.SendString("tablas creadas porfavorsss")
	})

	app.Listen(":8080")
}