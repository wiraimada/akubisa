package routes

import (
	"akubisa/internal/ai"
	"akubisa/internal/auth"
	"akubisa/internal/lessons" // Import the lessons package
	"akubisa/internal/shared/storage"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm" // Import gorm
)

// Register sets up all the application routes
func Register(app *fiber.App, db *gorm.DB, jwtSecret string) { // Change db type to *gorm.DB
	// Auth module setup
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo, jwtSecret)
	authHandler := auth.NewHandler(authService)

	// Lessons module setup
	lessonRepo := lessons.NewRepository(db)
	lessonService := lessons.NewService(lessonRepo)
	lessonHandler := lessons.NewHandler(lessonService)

	// AI module setup
	aiRepo := ai.NewRepository(db)
	openaiClient := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	s3Storage, _ := storage.NewS3Storage(
		os.Getenv("S3_REGION"),
		os.Getenv("S3_ENDPOINT"),
		os.Getenv("S3_ACCESS_KEY"),
		os.Getenv("S3_SECRET_KEY"),
		os.Getenv("S3_USE_PATH_STYLE") == "true",
	)

	aiService := ai.NewService(aiRepo, openaiClient, s3Storage, os.Getenv("S3_BUCKET"))
	aiHandler := ai.NewHandler(aiService)

	api := app.Group("/api/v1")

	// Auth routes
	api.Post("/auth/register", authHandler.Register)
	api.Post("/auth/login", authHandler.Login)

	// Lesson routes
	api.Get("/lessons", lessonHandler.ListLessons)   // Corrected handler method
	api.Get("/lessons/:id", lessonHandler.GetLesson) // Corrected handler method

	// Quiz submission route
	api.Post("/quizzes/:id/submit", lessonHandler.SubmitQuiz)

	// AI routes
	api.Post("/ai/pronunciation/analyze", aiHandler.AnalyzePronunciation)
}
