package main

import (
	"log"
	"os"

	"akubisa/internal/routes"
	"akubisa/internal/shared/database"
	"akubisa/internal/shared/database/seeders" // Import the seeders package
	"akubisa/pkg/logger"
	"akubisa/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	logger.Init(os.Getenv("APP_ENV"))

	db, err := database.NewPostgresConnection(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	if err != nil {
		log.Fatal(err)
	}

	// Run seeders
	seeders.SeedLessons(db)

	app := fiber.New()

	app.Use(requestid.New())

	app.Use(func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				logger.Logger.Error("panic_recovered",
					"error", r,
					"request_id", c.Get("X-Request-ID"),
				)
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"success": false,
					"message": "internal server error",
				})
			}
		}()
		return c.Next()
	})

	app.Use(middleware.AccessLogger())

	routes.Register(app, db)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "API Running",
		})
	})

	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}
