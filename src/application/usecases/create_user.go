package usecases

import (
	"github.com/google/uuid"
	"github.com/julianojj/cop_golang/src/domain/entities"
	"github.com/julianojj/cop_golang/src/domain/repositories"
)

type CreateUser struct {
	UserRepository repositories.UserRepository
}

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
}

func NewCreateUser(
	userRepository repositories.UserRepository,
) *CreateUser {
	return &CreateUser{
		UserRepository: userRepository,
	}
}

func (createUser *CreateUser) Execute(input CreateUserInput) {
	user, _ := entities.NewUser(
		uuid.NewString(),
		input.Name,
		input.Email,
		input.Password,
	)
	createUser.UserRepository.Save(user)
}
