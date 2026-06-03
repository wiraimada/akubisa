package auth

import (
	"errors"
	"os"

	"golang.org/x/crypto/bcrypt"

	"akubisa/internal/shared/utils"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(req RegisterRequest) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return err
	}

	user := User{
		FullName:     req.FullName,
		Email:        req.Email,
		PasswordHash: string(hashed),
		Role:         "parent",
	}

	return s.repo.CreateUser(user)
}

func (s *Service) Login(req LoginRequest) (string, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(
		user.ID.String(),
		os.Getenv("JWT_SECRET"),
	)

	if err != nil {
		return "", err
	}

	return token, nil
}
