package user

import (
	"os/user"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) FindByID(id int) (*user.User, error) {
	args := m.Called(id)
	return args.Get(0).(*user.User), args.Error(1)
}
