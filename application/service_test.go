package application

import (
	"testing"

	domain "poc-injecao-dependencia/domain"
	infrastructure_user "poc-injecao-dependencia/infrastructure"

	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
)

func TestUserService_GetUserByID(t *testing.T) {
	container := dig.New()

	mockRepo := &infrastructure_user.MockRepository{}
	mockRepo.On("FindByID", 1).Return(&domain.User{ID: 1, Name: "John Doe", Email: "john@example.com"}, nil)

	if err := container.Provide(func() domain.Repository {
		return mockRepo
	}); err != nil {
		t.Fatal(err)
	}

	if err := container.Provide(NewService); err != nil {
		t.Fatal(err)
	}

	err := container.Invoke(func(service *Service) {
		// Agora o serviço já está configurado com o mockRepo
		user, err := service.GetUserByID(1)
		assert.NoError(t, err)
		assert.Equal(t, 1, user.ID)
		assert.Equal(t, "John Doe", user.Name)
		assert.Equal(t, "john@example.com", user.Email)

		mockRepo.AssertCalled(t, "FindByID", 1)
	})

	assert.NoError(t, err)
}
