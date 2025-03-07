package application

import "poc-injecao-dependencia/domain"

type Service struct {
	Repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) GetUserByID(id int) (*domain.User, error) {
	return s.Repo.FindByID(id)
}
