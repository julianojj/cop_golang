package integration_test

import (
	"testing"

	"github.com/julianojj/cop_golang/src/application/usecases"
	"github.com/julianojj/cop_golang/src/infra/repositories/memory"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	userRepository := memory.NewUserRepositoryMemory()
	createUser := usecases.NewCreateUser(userRepository)
	createUser.Execute(usecases.CreateUserInput{
		Name:     "Juliano",
		Email:    "juliano@test.com",
		Password: "123456",
	})
	users := userRepository.FindAll()
	assert.Equal(t, len(users), 1)
	userRepository.Clean()
}
