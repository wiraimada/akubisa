package lessons

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Lesson represents a learning lesson
type Lesson struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	ContentType string         `gorm:"type:varchar(50);not null" json:"content_type"` // e.g., "text", "video", "interactive"
	Content     string         `gorm:"type:text" json:"content"`                     // actual content or URL to content
	Order       int            `gorm:"type:integer;not null" json:"order"`
	CreatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Quizzes []Quiz `gorm:"foreignKey:LessonID" json:"quizzes"`
}

// Quiz represents a quiz associated with a lesson
type Quiz struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	LessonID    uuid.UUID      `gorm:"type:uuid;not null" json:"lesson_id"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	Order       int            `gorm:"type:integer;not null" json:"order"`
	CreatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Questions []QuizQuestion `gorm:"foreignKey:QuizID" json:"questions"`
}

// QuizQuestion represents a question in a quiz
type QuizQuestion struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	QuizID      uuid.UUID      `gorm:"type:uuid;not null" json:"quiz_id"`
	QuestionText string         `gorm:"type:text;not null" json:"question_text"`
	QuestionType string         `gorm:"type:varchar(50);not null" json:"question_type"` // e.g., "multiple_choice", "fill_in_the_blank"
	Order       int            `gorm:"type:integer;not null" json:"order"`
	CreatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Answers []QuizAnswer `gorm:"foreignKey:QuestionID" json:"answers"`
}

// QuizAnswer represents an answer option for a quiz question
type QuizAnswer struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	QuestionID  uuid.UUID      `gorm:"type:uuid;not null" json:"question_id"`
	AnswerText  string         `gorm:"type:text;not null" json:"answer_text"`
	IsCorrect   bool           `gorm:"type:boolean;default:false" json:"is_correct"`
	Order       int            `gorm:"type:integer;not null" json:"order"`
	CreatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// Progress represents a child's progress on a lesson or quiz
type Progress struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ChildID     uuid.UUID      `gorm:"type:uuid;not null" json:"child_id"` // Assuming a Child model exists or will exist
	LessonID    *uuid.UUID     `gorm:"type:uuid" json:"lesson_id,omitempty"`
	QuizID      *uuid.UUID     `gorm:"type:uuid" json:"quiz_id,omitempty"`
	Score       *int           `gorm:"type:integer" json:"score,omitempty"`
	Completed   bool           `gorm:"type:boolean;default:false" json:"completed"`
	CompletedAt *time.Time     `json:"completed_at,omitempty"`
	CreatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
