package repositories

import "github.com/julianojj/cop_golang/src/domain/entities"

type UserRepository interface {
	Save(user *entities.User)
	FindAll() []*entities.User
	Clean()
}
