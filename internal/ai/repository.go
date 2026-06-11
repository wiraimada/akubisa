package ai

import (
	"gorm.io/gorm"
)

// Repository defines the interface for AI data operations
type Repository interface {
	SaveResult(result *PronunciationResult) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new AI repository
func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) SaveResult(result *PronunciationResult) error {
	return r.db.Create(result).Error
}
