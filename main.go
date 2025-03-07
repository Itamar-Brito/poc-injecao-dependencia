package main

import (
	"fmt"
	"net/http"
	"strconv"

	"poc-injecao-dependencia/application"
	domain "poc-injecao-dependencia/domain"
	infrastructure "poc-injecao-dependencia/infrastructure/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func setupRouter(container *dig.Container) *gin.Engine {
	router := gin.Default()

	router.GET("/users/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		err = container.Invoke(func(service *application.Service) {
			user, err := service.GetUserByID(id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			if user == nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}

			c.JSON(http.StatusOK, user)
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Container error: " + err.Error()})
		}
	})

	return router
}

func main() {
	container := BuildContainer()

	if err := container.Provide(func() domain.Repository {
		// You would replace this with a real repository implementation
		return &infrastructure.RealRepository{}
	}, dig.Name("real")); err != nil {
		fmt.Printf("Failed to provide real repository: %v\n", err)
	}

	router := setupRouter(container)
	router.Run(":8080")
}

func BuildContainer() *dig.Container {
	container := dig.New()

	// Register real repository
	container.Provide(infrastructure.NewRealRepository)

	// Register the service
	container.Provide(application.NewService)

	return container
}
