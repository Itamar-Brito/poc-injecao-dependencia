package user

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetUserByID(id int) (*User, error) {
	return s.repo.FindByID(id)
}
