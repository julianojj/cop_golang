package entities

import (
	"errors"
	"net/mail"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(
	id string,
	name string,
	email string,
	password string,
) (*User, error) {
	if isInvalidEmail(email) {
		return nil, errors.New("invalid email")
	}
	if isInvalidPassword(password) {
		return nil, errors.New("invalid password")
	}
	return &User{
		Id:       id,
		Name:     name,
		Email:    password,
		Password: password,
	}, nil
}

func isInvalidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err != nil

}

func isInvalidPassword(password string) bool {
	return len(password) < 6
}
