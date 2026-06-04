package lessons

import (
	"time"

	"github.com/google/uuid"
)

type Lesson struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Difficulty  int       `json:"difficulty"`
	Content     any       `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
}
