package seeders

import (
	"log"
	"time"

	"akubisa/internal/lessons" // Import the lessons package to use Lesson, Quiz, etc.
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SeedLessons seeds initial lesson data into the database
func SeedLessons(db *gorm.DB) {
	log.Println("Seeding lessons data...")

	lessonsData := []lessons.Lesson{
		{
			Title:       "Membaca Dasar: Mengenal Huruf A-Z",
			Description: "Pelajari huruf-huruf dasar dari A sampai Z dan cara membacanya.",
			ContentType: "text",
			Content:     "Konten pelajaran membaca dasar...",
			Order:       1,
			Quizzes: []lessons.Quiz{
				{
					Title:       "Kuis Huruf A-Z",
					Description: "Uji kemampuanmu mengenal huruf A-Z.",
					Order:       1,
					Questions: []lessons.QuizQuestion{
						{
							QuestionText: "Huruf apakah ini: A?",
							QuestionType: "multiple_choice",
							Order:        1,
							Answers: []lessons.QuizAnswer{
								{AnswerText: "A", IsCorrect: true, Order: 1},
								{AnswerText: "B", IsCorrect: false, Order: 2},
								{AnswerText: "C", IsCorrect: false, Order: 3},
							},
						},
						{
							QuestionText: "Pilih huruf yang benar untuk 'Buku'",
							QuestionType: "multiple_choice",
							Order:        2,
							Answers: []lessons.QuizAnswer{
								{AnswerText: "A", IsCorrect: false, Order: 1},
								{AnswerText: "B", IsCorrect: true, Order: 2},
								{AnswerText: "C", IsCorrect: false, Order: 3},
							},
						},
					},
				},
			},
		},
		{
			Title:       "Menulis Dasar: Menulis Huruf A-Z",
			Description: "Latihan menulis huruf-huruf dasar dari A sampai Z.",
			ContentType: "interactive",
			Content:     "Konten pelajaran menulis dasar...",
			Order:       2,
			Quizzes: []lessons.Quiz{
				{
					Title:       "Kuis Menulis Huruf A-Z",
					Description: "Uji kemampuanmu menulis huruf A-Z.",
					Order:       1,
					Questions: []lessons.QuizQuestion{
						{
							QuestionText: "Bagaimana cara menulis huruf 'D'?",
							QuestionType: "fill_in_the_blank",
							Order:        1,
							Answers: []lessons.QuizAnswer{
								{AnswerText: "D", IsCorrect: true, Order: 1},
							},
						},
					},
				},
			},
		},
		{
			Title:       "Berhitung Dasar: Angka 1-10",
			Description: "Pelajari angka-angka dasar dari 1 sampai 10 dan cara menghitungnya.",
			ContentType: "text",
			Content:     "Konten pelajaran berhitung dasar...",
			Order:       3,
			Quizzes: []lessons.Quiz{
				{
					Title:       "Kuis Angka 1-10",
					Description: "Uji kemampuanmu mengenal angka 1-10.",
					Order:       1,
					Questions: []lessons.QuizQuestion{
						{
							QuestionText: "Berapakah 1 + 1?",
							QuestionType: "multiple_choice",
							Order:        1,
							Answers: []lessons.QuizAnswer{
								{AnswerText: "1", IsCorrect: false, Order: 1},
								{AnswerText: "2", IsCorrect: true, Order: 2},
								{AnswerText: "3", IsCorrect: false, Order: 3},
							},
						},
						{
							QuestionText: "Angka berapa setelah 5?",
							QuestionType: "multiple_choice",
							Order:        2,
							Answers: []lessons.QuizAnswer{
								{AnswerText: "4", IsCorrect: false, Order: 1},
								{AnswerText: "6", IsCorrect: true, Order: 2},
								{AnswerText: "7", IsCorrect: false, Order: 3},
							},
						},
					},
				},
			},
		},
	}

	for _, lesson := range lessonsData {
		// Check if lesson already exists to prevent duplicates
		var existingLesson lessons.Lesson
		if err := db.Where("title = ?", lesson.Title).First(&existingLesson).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Lesson does not exist, create it
				if err := db.Create(&lesson).Error; err != nil {
					log.Printf("Failed to seed lesson '%s': %v", lesson.Title, err)
				} else {
					log.Printf("Seeded lesson: %s", lesson.Title)
				}
			} else {
				log.Printf("Error checking for existing lesson '%s': %v", lesson.Title, err)
			}
		} else {
			log.Printf("Lesson '%s' already exists, skipping.", lesson.Title)
		}
	}
	log.Println("Lessons seeding completed.")
}

// BeforeCreate hook to generate UUIDs
func (l *lessons.Lesson) BeforeCreate(tx *gorm.DB) (err error) {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}
	l.CreatedAt = time.Now()
	l.UpdatedAt = time.Now()
	return
}

func (q *lessons.Quiz) BeforeCreate(tx *gorm.DB) (err error) {
	if q.ID == uuid.Nil {
		q.ID = uuid.New()
	}
	q.CreatedAt = time.Now()
	q.UpdatedAt = time.Now()
	return
}

func (qq *lessons.QuizQuestion) BeforeCreate(tx *gorm.DB) (err error) {
	if qq.ID == uuid.Nil {
		qq.ID = uuid.New()
	}
	qq.CreatedAt = time.Now()
	qq.UpdatedAt = time.Now()
	return
}

func (qa *lessons.QuizAnswer) BeforeCreate(tx *gorm.DB) (err error) {
	if qa.ID == uuid.Nil {
		qa.ID = uuid.New()
	}
	qa.CreatedAt = time.Now()
	qa.UpdatedAt = time.Now()
	return
}

func (p *lessons.Progress) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return
}

// BeforeUpdate hook to update UpdatedAt
func (l *lessons.Lesson) BeforeUpdate(tx *gorm.DB) (err error) {
	l.UpdatedAt = time.Now()
	return
}

func (q *lessons.Quiz) BeforeUpdate(tx *gorm.DB) (err error) {
	q.UpdatedAt = time.Now()
	return
}

func (qq *lessons.QuizQuestion) BeforeUpdate(tx *gorm.DB) (err error) {
	qq.UpdatedAt = time.Now()
	return
}

func (qa *lessons.QuizAnswer) BeforeUpdate(tx *gorm.DB) (err error) {
	qa.UpdatedAt = time.Now()
	return
}

func (p *lessons.Progress) BeforeUpdate(tx *gorm.DB) (err error) {
	p.UpdatedAt = time.Now()
	return
}
