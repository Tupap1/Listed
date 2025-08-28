package main

import (
		"fmt"
		"github.com/gofiber/fiber/v2"
		"gorm.io/driver/mysql"
		"gorm.io/gorm"
		"os"
		"log"
		"github.com/Tupap1/Listed/server/models"
		"github.com/joho/godotenv"
	)

func main() {

	err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
  		}
	dsn := os.Getenv("LOCALDB_DSN")
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database:", err)
	}


	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Andres eres el mejor del mundo")
	})
	
	app.Get("/db/update", func(c *fiber.Ctx) error {
		db.AutoMigrate(&models.RefreshToken{})
		return c.SendString("Update DB")
	})

	app.Listen(":8080")
}