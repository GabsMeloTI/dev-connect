package service

import (
	"context"
	"database/sql"
	"errors"
	"treads/internal/model"
	"treads/internal/repository"
	"treads/pkg/crypt"
)

type UserInterface interface {
	CreateUser(context.Context, model.UserCreateDto) (model.UserResponse, error)
	UpdateUser(context.Context, model.UserUpdateDto) (model.UserResponse, error)
	DeleteUser(context.Context, int64) error
	GetAllUsers(context.Context) ([]model.UserResponse, error)
}

type User struct {
	UserInterface repository.UserInterface
}

func NewUser(UserInterface repository.UserInterface) *User {
	return &User{UserInterface: UserInterface}
}

func (s *User) CreateUser(ctx context.Context, data model.UserCreateDto) (model.UserResponse, error) {
	arg := data.ParseCreateToUser()

	existsUsername, err := s.UserInterface.GetUsersByName(ctx, data.Username)
	if existsUsername {
		return model.UserResponse{}, errors.New("Username already exists")
	}
	if err != nil {
		return model.UserResponse{}, err
	}

	existsEmail, err := s.UserInterface.GetUsersByEmail(ctx, data.Email)
	if existsEmail {
		return model.UserResponse{}, errors.New("Email already exists")
	}
	if err != nil {
		return model.UserResponse{}, err
	}

	EncryptedPass, err := crypt.HashPassword(data.Password)
	if err != nil {
		return model.UserResponse{}, err
	}
	result, err := s.UserInterface.CreateUser(ctx, arg)
	result.Password = EncryptedPass
	if err != nil {
		return model.UserResponse{}, err
	}

	createUser := model.UserResponse{}
	createUser.ParseFromUserObject(result)

	return createUser, nil
}

func (s *User) UpdateUser(ctx context.Context, data model.UserUpdateDto) (model.UserResponse, error) {
	arg := data.ParseUpdateToUser()

	getResult, err := s.UserInterface.GetUserById(ctx, data.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.UserResponse{}, errors.New("User não encontrado")
		}
		return model.UserResponse{}, err
	}

	if getResult.Username != data.Username || getResult.Email != data.Email {
		existsUsername, err := s.UserInterface.GetUsersByName(ctx, data.Username)
		if existsUsername {
			return model.UserResponse{}, errors.New("Username already exists")
		}
		if err != nil {
			return model.UserResponse{}, err
		}

		existsEmail, err := s.UserInterface.GetUsersByEmail(ctx, data.Email)
		if existsEmail {
			return model.UserResponse{}, errors.New("Email already exists")
		}
		if err != nil {
			return model.UserResponse{}, err
		}
	}

	EncryptedPass, err := crypt.HashPassword(data.Password)
	result, err := s.UserInterface.UpdateUser(ctx, arg)
	result.Password = EncryptedPass
	if err != nil {
		return model.UserResponse{}, err
	}

	updateUser := model.UserResponse{}
	updateUser.ParseFromUserObject(result)

	return updateUser, nil
}

func (s *User) DeleteUser(ctx context.Context, id int64) error {
	_, err := s.UserInterface.GetUserById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("User não encontrado")
		}
		return err
	}

	return s.UserInterface.DeleteUser(ctx, id)
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
