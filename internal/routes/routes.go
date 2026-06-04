package routes

import (
	"akubisa/internal/auth"
	"akubisa/internal/lessons" // Import the lessons package
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm" // Import gorm
)

// Register sets up all the application routes
func Register(app *fiber.App, db *gorm.DB) { // Change db type to *gorm.DB
	// Auth module setup
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	// Lessons module setup
	lessonRepo := lessons.NewRepository(db)
	lessonService := lessons.NewService(lessonRepo)
	lessonHandler := lessons.NewHandler(lessonService)

	api := app.Group("/api/v1")

	// Auth routes
	api.Post("/auth/register", authHandler.Register)
	api.Post("/auth/login", authHandler.Login)

	// Lesson routes
	api.Get("/lessons", lessonHandler.ListLessons)   // Corrected handler method
	api.Get("/lessons/:id", lessonHandler.GetLesson) // Corrected handler method

	// Quiz submission route
	api.Post("/quizzes/:id/submit", lessonHandler.SubmitQuiz)
}
