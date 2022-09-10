package unit_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/julianojj/cop_golang/src/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestUserWithInvalidEmail(t *testing.T) {
	_, err := entities.NewUser(
		uuid.NewString(),
		"Juliano",
		"julianotest.com",
		"123456",
	)
	assert.Equal(t, err.Error(), "invalid email")
}

func TestUserWithInvalidPassword(t *testing.T) {
	_, err := entities.NewUser(
		uuid.NewString(),
		"Juliano",
		"juliano@test.com",
		"12345",
	)
	assert.Equal(t, err.Error(), "invalid password")
}

func TestUser(t *testing.T) {
	user, _ := entities.NewUser(
		uuid.NewString(),
		"Juliano",
		"juliano@test.com",
		"123456",
	)
	assert.Equal(t, user.Name, "Juliano")
}
