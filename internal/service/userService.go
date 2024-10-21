package service

import "treads/internal/repository"

type User struct {
	UserInterface repository.UserInterface
}

func NewUser(UserInterface repository.UserInterface) *User {
	return &User{UserInterface: UserInterface}
}
