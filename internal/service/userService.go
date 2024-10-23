package service

import (
	"context"
	"treads/internal/model"
	"treads/internal/repository"
)

type UserInterface interface {
	//CreateUser(context.Context, model.UserCreateDto) (model.UserRespose, error)
	//UpdateUser(context.Context, model.UserUpdateDto) (model.UserRespose, error)
	//DeleteUser(context.Context, int32) error
	GetAllUsers(context.Context) ([]model.UserResponse, error)
}

type User struct {
	UserInterface repository.UserInterface
}

func NewUser(UserInterface repository.UserInterface) *User {
	return &User{UserInterface: UserInterface}
}

func (s *User) GetAllUsers(ctx context.Context) ([]model.UserResponse, error) {
	results, err := s.UserInterface.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	getAllUser := model.UserResponse{}
	var usersResponse []model.UserResponse
	for _, result := range results {
		getAllUser.ParseFromUserObject(result)
		usersResponse = append(usersResponse, getAllUser)
	}

	return usersResponse, nil
}
