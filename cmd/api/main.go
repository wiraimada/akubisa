package main

import (
	"akubisa/pkg/database"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize Database
	database.InitDB()

	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Akubisa API!")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "UP",
		})
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
