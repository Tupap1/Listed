package main

import (
	"log"
	"os"

	"auth_service/handlers"
	"auth_service/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"

)

func main() {
	err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
  		}

	dsn := os.Getenv("LOCALDB_DSN")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	app := fiber.New()

	authService := services.NewAuthService(db)
	authHandler := handlers.NewAuthHandler(authService)

	authRoutes := app.Group("/api/v1/auth")
	{
		authRoutes.Post("/register", authHandler.Register)
		authRoutes.Post("/login", authHandler.Login)
		authRoutes.Post("/refresh", authHandler.Refresh)
		authRoutes.Post("/logout", authHandler.Logout)

	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	
	log.Printf("Authentication service is running on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
