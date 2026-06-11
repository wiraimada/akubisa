package ai

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PronunciationResult represents the result of an AI pronunciation analysis
type PronunciationResult struct {
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	ChildID       uuid.UUID      `gorm:"type:uuid;not null" json:"child_id"`
	AudioURL      string         `gorm:"type:text;not null" json:"audio_url"`
	Transcription string         `gorm:"type:text" json:"transcription"`
	Score         float64        `gorm:"type:decimal(5,2)" json:"score"`
	Feedback      string         `gorm:"type:text" json:"feedback"`
	CreatedAt     time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (p *PronunciationResult) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return
}
