package database

import (
	"fmt"
	"log"
	"os"

	"akubisa/internal/auth"
	"akubisa/internal/lessons" // Import lessons models

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection(host, port, user, password, dbname string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host,
		user,
		password,
		dbname,
		port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Enable the uuid-ossp extension so uuid_generate_v4() column defaults resolve.
	if err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		return nil, fmt.Errorf("failed to enable uuid-ossp extension: %w", err)
	}

	// AutoMigrate all models
	err = db.AutoMigrate(
		&auth.User{},
		&lessons.Lesson{},
		&lessons.Quiz{},
		&lessons.QuizQuestion{},
		&lessons.QuizAnswer{},
		&lessons.Progress{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate database: %w", err)
	}

	log.Println("Database connection and auto-migration successful!")

	return db, nil
}

// GetDB returns the GORM DB instance. This is a placeholder for now,
// as the actual DB instance will be passed around or managed by a dependency injection.
// For simplicity, we'll keep it as a global or pass it from main.
// This function might be removed or refactored later.
func GetDB() *gorm.DB {
	// This is a temporary solution. In a real application,
	// you would use dependency injection or pass the DB instance.
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err := NewPostgresConnection(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Failed to get DB instance: %v", err)
	}
	return db
}
