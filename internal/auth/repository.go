package auth

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user User) error
	FindByEmail(email string) (*User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(user User) error {
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	return r.db.Create(&user).Error
}

func (r *repository) FindByEmail(email string) (*User, error) {
	var user User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
