package memory

import "github.com/julianojj/cop_golang/src/domain/entities"

var users []*entities.User

type UserRepositoryMemory struct{}

func NewUserRepositoryMemory() *UserRepositoryMemory {
	return &UserRepositoryMemory{}
}

func (*UserRepositoryMemory) Save(user *entities.User) {
	users = append(users, user)
}

func (*UserRepositoryMemory) FindAll() []*entities.User {
	return users
}

func (*UserRepositoryMemory) Clean() {
	users = []*entities.User{}
}
