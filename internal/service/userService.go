package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"
	"treads/internal/helper"
	"treads/internal/model"
	"treads/internal/repository"
	"treads/pkg/crypt"
	"treads/pkg/token"
)

type UserInterface interface {
	CreateUser(context.Context, model.UserCreateDto) (model.UserResponse, error)
	UpdateUser(context.Context, model.UserUpdateDto) (model.UserResponse, error)
	UpdatePassword(context.Context, model.UserRequestUpdatePasswordByUser) error
	DisableUser(context.Context, int64) error
	DeleteUser(context.Context, int64) error
	GetAllUsers(context.Context, string) ([]model.UserResponse, error)
	UserLogin(context.Context, model.LoginUserRequest) (model.LoginUserResponse, error)
}

type User struct {
	UserInterface repository.UserInterface
}

func NewUser(UserInterface repository.UserInterface) *User {
	return &User{UserInterface: UserInterface}
}

func (s *User) CreateUser(ctx context.Context, data model.UserCreateDto) (model.UserResponse, error) {
	exists, err := s.UserInterface.GetUsersByUsernameOrEmail(ctx, data.Username)
	if err != nil {
		return model.UserResponse{}, err
	}
	if exists {
		return model.UserResponse{}, errors.New("username or email already exists")
	}

	hashedPassword, err := crypt.HashPassword(data.Password)
	if err != nil {
		return model.UserResponse{}, err
	}
	data.Password = hashedPassword

	arg := data.ParseCreateToUser()
	result, err := s.UserInterface.CreateUser(ctx, arg)
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
			return model.UserResponse{}, errors.New("user not found")
		}
		return model.UserResponse{}, err
	}

	if getResult.Username != data.Username || getResult.Email != data.Email {
		exists, err := s.UserInterface.GetUsersByUsernameOrEmail(ctx, data.Username)
		if err != nil {
			return model.UserResponse{}, err
		}
		if exists {
			return model.UserResponse{}, errors.New("username or email already exists")
		}
	}

	result, err := s.UserInterface.UpdateUser(ctx, arg)
	if err != nil {
		return model.UserResponse{}, err
	}

	updateUser := model.UserResponse{}
	updateUser.ParseFromUserObject(result)

	return updateUser, nil
}

func (s *User) UpdatePassword(ctx context.Context, data model.UserRequestUpdatePasswordByUser) error {
	if data.Password != data.ConfirmPassword {
		return errors.New("passwords do not match")
	}

	encryptedPass, err := crypt.HashPassword(data.Password)
	if err != nil {
		return err
	}

	arg := data.ParseUpdateToPassword()
	arg.Password = encryptedPass

	return s.UserInterface.UpdatePassword(ctx, arg)
}

func (s *User) DisableUser(ctx context.Context, id int64) error {
	_, err := s.UserInterface.GetUserById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("user not found")
		}
		return err
	}

	return s.UserInterface.DisableUser(ctx, id)
}

func (s *User) DeleteUser(ctx context.Context, id int64) error {
	_, err := s.UserInterface.GetUserById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("user not found")
		}
		return err
	}

	return s.UserInterface.DeleteUser(ctx, id)
}

func (s *User) GetAllUsers(ctx context.Context, username string) ([]model.UserResponse, error) {
	arg := sql.NullString{String: username, Valid: username != ""}

	results, err := s.UserInterface.GetAllUsers(ctx, arg)
	if err != nil {
		return nil, err
	}

	var usersResponse []model.UserResponse
	for _, result := range results {
		user := model.UserResponse{}
		user.ParseFromUserObject(result)
		usersResponse = append(usersResponse, user)
	}

	return usersResponse, nil
}

func (s *User) UserLogin(ctx context.Context, data model.LoginUserRequest) (model.LoginUserResponse, error) {
	result, err := s.UserInterface.GetUsersLoginByEmailOrUsername(ctx, data.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.LoginUserResponse{}, errors.New("invalid credentials")
		}
		return model.LoginUserResponse{}, err
	}

	if !crypt.CheckPasswordHash(data.Password, result.Password) {
		return model.LoginUserResponse{}, errors.New("invalid credentials")
	}

	symetricKey := helper.GetSignatureString()

	maker, err := token.NewPasetoMaker(symetricKey)
	if err != nil {
		return model.LoginUserResponse{}, fmt.Errorf("token creation error: %w", err)
	}

	idStr := strconv.FormatInt(result.ID, 10)
	tokenStr, err := maker.CreateToken(idStr, result.Username, result.Name, result.Email, 24*time.Hour)
	if err != nil {
		return model.LoginUserResponse{}, fmt.Errorf("token generation failed: %w", err)
	}

	return model.LoginUserResponse{
		ID:        idStr,
		Name:      result.Name,
		Email:     result.Email,
		Username:  result.Username,
		Token:     tokenStr,
		LastLogin: time.Now(),
	}, nil
}
