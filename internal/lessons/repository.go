package lessons

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Repository defines the interface for lesson data operations
type Repository interface {
	GetAllLessons() ([]Lesson, error)
	GetLessonByID(id uuid.UUID) (*Lesson, error)
	SubmitQuiz(quizID uuid.UUID, childID uuid.UUID, answers []QuizSubmission) (*Progress, error) // Modified return type
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new lesson repository
func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAllLessons() ([]Lesson, error) {
	var lessons []Lesson
	err := r.db.Preload("Quizzes.Questions.Answers").Order("order ASC").Find(&lessons).Error
	return lessons, err
}

func (r *repository) GetLessonByID(id uuid.UUID) (*Lesson, error) {
	var lesson Lesson
	err := r.db.Preload("Quizzes.Questions.Answers").First(&lesson, "id = ?", id).Error
	return &lesson, err
}

// QuizSubmission represents a user's answer to a quiz question
type QuizSubmission struct {
	QuestionID uuid.UUID
	AnswerIDs  []uuid.UUID // For multiple choice, can be multiple. For single choice, just one.
	AnswerText string      // For fill-in-the-blank
}

func (r *repository) SubmitQuiz(quizID uuid.UUID, childID uuid.UUID, submissions []QuizSubmission) (*Progress, error) { // Modified return type
	// Start a transaction
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	var quiz Quiz
	if err := tx.Preload("Questions.Answers").First(&quiz, "id = ?", quizID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	correctAnswersCount := 0
	totalQuestions := len(quiz.Questions)

	for _, submission := range submissions {
		var question QuizQuestion
		if err := tx.Preload("Answers").First(&question, "id = ?", submission.QuestionID).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		isCorrect := false
		if question.QuestionType == "multiple_choice" {
			// Check if all submitted answers are correct and no incorrect answers are selected
			correctSubmissionCount := 0
			for _, submittedAnswerID := range submission.AnswerIDs {
				for _, qa := range question.Answers {
					if qa.ID == submittedAnswerID && qa.IsCorrect {
						correctSubmissionCount++
						break
					}
				}
			}

			// Count total correct answers for the question
			totalCorrectAnswersForQuestion := 0
			for _, qa := range question.Answers {
				if qa.IsCorrect {
					totalCorrectAnswersForQuestion++
				}
			}

			if correctSubmissionCount == totalCorrectAnswersForQuestion && len(submission.AnswerIDs) == totalCorrectAnswersForQuestion {
				isCorrect = true
			}
		} else if question.QuestionType == "fill_in_the_blank" {
			// For fill-in-the-blank, check if the submitted text matches any correct answer
			for _, qa := range question.Answers {
				if qa.IsCorrect && qa.AnswerText == submission.AnswerText {
					isCorrect = true
					break
				}
			}
		}

		if isCorrect {
			correctAnswersCount++
		}
	}

	score := (correctAnswersCount * 100) / totalQuestions
	completed := false
	var completedAt *time.Time
	if score >= 70 { // Assuming 70% is the passing score
		completed = true
		now := time.Now()
		completedAt = &now
	}

	progress := Progress{
		ChildID:     childID,
		QuizID:      &quizID,
		Score:       &score,
		Completed:   completed,
		CompletedAt: completedAt,
	}

	if err := tx.Create(&progress).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &progress, tx.Commit().Error // Return the created progress
}
