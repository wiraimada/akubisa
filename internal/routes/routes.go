package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"

	"akubisa/internal/auth"
	"akubisa/internal/lessons"
)

func Register(app *fiber.App, db *pgxpool.Pool) {
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	lessonRepo := lessons.NewRepository(db)
	lessonService := lessons.NewService(lessonRepo)
	lessonHandler := lessons.NewHandler(lessonService)

	api := app.Group("/api/v1")

	api.Post("/auth/register", authHandler.Register)
	api.Post("/auth/login", authHandler.Login)

	api.Get("/lessons", lessonHandler.List)
	api.Get("/lessons/:id", lessonHandler.GetByID)
}
