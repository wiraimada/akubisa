package lessons

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BeforeCreate hooks generate UUIDs and set timestamps before insert.

func (l *Lesson) BeforeCreate(tx *gorm.DB) (err error) {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}
	l.CreatedAt = time.Now()
	l.UpdatedAt = time.Now()
	return
}

func (q *Quiz) BeforeCreate(tx *gorm.DB) (err error) {
	if q.ID == uuid.Nil {
		q.ID = uuid.New()
	}
	q.CreatedAt = time.Now()
	q.UpdatedAt = time.Now()
	return
}

func (qq *QuizQuestion) BeforeCreate(tx *gorm.DB) (err error) {
	if qq.ID == uuid.Nil {
		qq.ID = uuid.New()
	}
	qq.CreatedAt = time.Now()
	qq.UpdatedAt = time.Now()
	return
}

func (qa *QuizAnswer) BeforeCreate(tx *gorm.DB) (err error) {
	if qa.ID == uuid.Nil {
		qa.ID = uuid.New()
	}
	qa.CreatedAt = time.Now()
	qa.UpdatedAt = time.Now()
	return
}

func (p *Progress) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return
}

// BeforeUpdate hooks refresh the UpdatedAt timestamp.

func (l *Lesson) BeforeUpdate(tx *gorm.DB) (err error) {
	l.UpdatedAt = time.Now()
	return
}

func (q *Quiz) BeforeUpdate(tx *gorm.DB) (err error) {
	q.UpdatedAt = time.Now()
	return
}

func (qq *QuizQuestion) BeforeUpdate(tx *gorm.DB) (err error) {
	qq.UpdatedAt = time.Now()
	return
}

func (qa *QuizAnswer) BeforeUpdate(tx *gorm.DB) (err error) {
	qa.UpdatedAt = time.Now()
	return
}

func (p *Progress) BeforeUpdate(tx *gorm.DB) (err error) {
	p.UpdatedAt = time.Now()
	return
}
