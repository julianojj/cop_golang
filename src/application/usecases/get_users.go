package usecases

import "github.com/julianojj/cop_golang/src/domain/repositories"

type GetUsers struct {
	UserRepository repositories.UserRepository
}

type getUsersOutput struct {
	id    string
	name  string
	email string
}

func NewGetUsers(
	userRepository repositories.UserRepository,
) *GetUsers {
	return &GetUsers{
		UserRepository: userRepository,
	}
}

func (getUsers *GetUsers) Execute() []getUsersOutput {
	users := getUsers.UserRepository.FindAll()
	var output []getUsersOutput
	for _, user := range users {
		output = append(output, getUsersOutput{
			id:    user.Id,
			name:  user.Name,
			email: user.Email,
		})
	}
	return output
}
