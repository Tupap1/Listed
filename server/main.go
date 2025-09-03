package main

import (
		"github.com/gofiber/fiber/v2"
		"gorm.io/driver/mysql"
		"gorm.io/gorm"
		"os"
		"log"
		"github.com/Tupap1/Listed/server/models"
		"github.com/Tupap1/Listed/server/database"
		"github.com/gofiber/fiber/v2/middleware/cors"
		"github.com/gofiber/fiber/v2/middleware/logger"
	)




func main() {
	db := connectDB()
	
	migrateDB(db)
	
	if shouldSeed() {
		if err := database.SeedDatabase(db); err != nil {
			log.Fatal("Error en seeding:", err)
		}
	}
	
	app := setupFiber()
	
	setupRoutes(app, db)
	
	log.Println("Servidor iniciado en puerto 8080")
	log.Fatal(app.Listen(":8080"))
}

func connectDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_DSN")
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}
	
	log.Println("✅ Conectado a MySQL")
	return db
}

func migrateDB(db *gorm.DB) {
	log.Println("🔄 Ejecutando migraciones...")
	
	err := db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.UserRole{},
		&models.RolePermission{},
		&models.RefreshToken{},
	)
	
	if err != nil {
		log.Fatal("Error en migraciones:", err)
	}
	
	log.Println("✅ Migraciones completadas")
}

func shouldSeed() bool {
	// Ejecutar seeding solo si:
	// 1. Es primera vez (variable de entorno)
	// 2. O en desarrollo
	return os.Getenv("SEED_DB") == "true" || os.Getenv("ENV") == "development"
}

func setupFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Inventory Management API",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})
	
	// Middlewares
	app.Use(logger.New())
	app.Use(cors.New())
	
	return app
}

func setupRoutes(app *fiber.App, db *gorm.DB) {
	// Rutas básicas
	api := app.Group("/api/v1")
	
	// Health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"message": "Inventory API is running",
		})
	})

}
	