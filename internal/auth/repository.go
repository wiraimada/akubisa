package auth

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

func (r *Repository) CreateUser(user User) error {
	query := `
        INSERT INTO users(id, full_name, email, password_hash, role)
        VALUES($1,$2,$3,$4,$5)
    `

	_, err := r.db.Exec(
		context.Background(),
		query,
		uuid.New(),
		user.FullName,
		user.Email,
		user.PasswordHash,
		user.Role,
	)

	return err
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	query := `
        SELECT id, full_name, email, password_hash, role
        FROM users
        WHERE email=$1
    `

	row := r.db.QueryRow(context.Background(), query, email)

	var user User

	err := row.Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
