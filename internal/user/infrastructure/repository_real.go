package infrastructure

import (
	"errors"
	domain "poc-injecao-dependencia/internal/user/domain"
)

type RealRepository struct {
	// This would typically have database connection details
	users map[int]*domain.User
}

func NewRealRepository() *RealRepository {
	// In a real application, you would initialize database connections here
	// For this example, we'll use an in-memory map
	return &RealRepository{
		users: map[int]*domain.User{
			1: {ID: 1, Name: "John Doe", Email: "john@example.com"},
			2: {ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
		},
	}
}

func (r *RealRepository) FindByID(id int) (*domain.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
