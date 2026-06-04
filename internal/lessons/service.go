package lessons

import (
	"errors"

	"github.com/google/uuid"
)

// Service defines the interface for lesson business logic
type Service interface {
	ListLessons() ([]Lesson, error)
	GetLesson(id uuid.UUID) (*Lesson, error)
	SubmitQuiz(quizID uuid.UUID, childID uuid.UUID, submissions []QuizSubmission) (*Progress, error)
}

type service struct {
	repo Repository
}

// NewService creates a new lesson service
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) ListLessons() ([]Lesson, error) {
	return s.repo.GetAllLessons()
}

func (s *service) GetLesson(id uuid.UUID) (*Lesson, error) {
	lesson, err := s.repo.GetLessonByID(id)
	if err != nil {
		return nil, errors.New("lesson not found")
	}
	return lesson, nil
}

func (s *service) SubmitQuiz(quizID uuid.UUID, childID uuid.UUID, submissions []QuizSubmission) (*Progress, error) {
	// In a real application, you might want to add more validation here
	// e.g., check if childID exists, if quizID exists and belongs to a lesson, etc.

	progress, err := s.repo.SubmitQuiz(quizID, childID, submissions)
	if err != nil {
		return nil, err
	}

	return progress, nil
}
