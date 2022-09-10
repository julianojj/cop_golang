package integration_test

import (
	"testing"

	"github.com/julianojj/cop_golang/src/application/usecases"
	"github.com/julianojj/cop_golang/src/infra/repositories/memory"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	userRepository := memory.NewUserRepositoryMemory()
	createUser := usecases.NewCreateUser(userRepository)
	inputUser := usecases.CreateUserInput{
		Name:     "Juliano",
		Email:    "juliano@test.com",
		Password: "123456",
	}
	createUser.Execute(inputUser)
	createUser.Execute(inputUser)
	createUser.Execute(inputUser)
	getUsers := usecases.NewGetUsers(userRepository)
	users := getUsers.Execute()
	assert.Equal(t, len(users), 3)
	userRepository.Clean()
}
