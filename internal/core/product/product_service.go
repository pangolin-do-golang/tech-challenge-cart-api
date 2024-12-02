package product

import (
	"github.com/google/uuid"
)

type Service struct {
	repository Repository
}

func NewProductService(repo Repository) *Service {
	return &Service{
		repository: repo,
	}
}

func (s *Service) Search(search string, category string) (*[]Product, error) {
	return s.repository.Search(search, category)
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.repository.Delete(id)
}

func (s *Service) GetByID(id uuid.UUID) (*Product, error) {
	return s.repository.GetByID(id)
}
