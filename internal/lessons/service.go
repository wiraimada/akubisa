package lessons

import "github.com/google/uuid"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) List() ([]Lesson, error) {
	return s.repo.List()
}

func (s *Service) FindByID(id string) (*Lesson, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.repo.FindByID(uuidID)
}
