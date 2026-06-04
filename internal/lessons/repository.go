package lessons

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) List() ([]Lesson, error) {
	query := `SELECT id, title, description, category, difficulty, content, created_at FROM lessons`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []Lesson
	for rows.Next() {
		var l Lesson
		err := rows.Scan(&l.ID, &l.Title, &l.Description, &l.Category, &l.Difficulty, &l.Content, &l.CreatedAt)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, l)
	}

	return lessons, nil
}

func (r *Repository) FindByID(id uuid.UUID) (*Lesson, error) {
	query := `SELECT id, title, description, category, difficulty, content, created_at FROM lessons WHERE id=$1`
	row := r.db.QueryRow(context.Background(), query, id)

	var l Lesson
	err := row.Scan(&l.ID, &l.Title, &l.Description, &l.Category, &l.Difficulty, &l.Content, &l.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &l, nil
}
