package model

import (
	"database/sql"
	"time"
	db "treads/db/sqlc"
)

type UserResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Bio       string `json:"bio"`
	AvatarUrl string `json:"avatar_url"`
}

type UserCreateDto struct {
	Name      string `json:"name"  validate:"required,min=3"`
	Username  string `json:"username"  validate:"required,min=3"`
	Email     string `json:"email"  validate:"required,min=3"`
	Password  string `json:"password"  validate:"required,min=3"`
	Bio       string `json:"bio"`
	AvatarUrl string `json:"avatar_url"`
}

type UserUpdateDto struct {
	ID        int64  `json:"id" validate:"required"`
	Name      string `json:"name"  validate:"required,min=3"`
	Username  string `json:"username"  validate:"required,min=3"`
	Email     string `json:"email"  validate:"required,min=3"`
	Bio       string `json:"bio"`
	AvatarUrl string `json:"avatar_url"`
}

type LoginUserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Token     string    `json:"token"`
	LastLogin time.Time `json:"last_login"`
}

type LoginUserRequest struct {
	Password string `query:"password" json:"password,omitempty" validate:"required"`
	Email    string `query:"email" json:"email,omitempty" validate:"required"`
}

type UserRequestUpdatePasswordByUser struct {
	ID              int64  `json:"id" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
}

type UserDeleteDto struct {
	ID int64 `json:"id"`
}

func (p *UserCreateDto) ParseCreateToUser() db.CreateUserParams {
	arg := db.CreateUserParams{
		Name:     p.Name,
		Username: p.Username,
		Email:    p.Email,
		Password: p.Password,
		Bio: sql.NullString{
			String: p.Bio,
			Valid:  true,
		},
		AvatarUrl: sql.NullString{
			String: p.AvatarUrl,
			Valid:  true,
		},
	}
	return arg
}

func (p *UserUpdateDto) ParseUpdateToUser() db.UpdateUserParams {
	arg := db.UpdateUserParams{
		Name:     p.Name,
		Username: p.Username,
		Email:    p.Email,
		Bio: sql.NullString{
			String: p.Bio,
			Valid:  true,
		},
		AvatarUrl: sql.NullString{
			String: p.AvatarUrl,
			Valid:  true,
		},
		ID: p.ID,
	}
	return arg
}

func (p *UserRequestUpdatePasswordByUser) ParseUpdateToPassword() db.UpdatePasswordByUserIdParams {
	arg := db.UpdatePasswordByUserIdParams{
		ID:       p.ID,
		Password: p.Password,
	}
	return arg
}

func (p *UserResponse) ParseFromUserObject(result db.User) {
	p.ID = result.ID
	p.Name = result.Name
	p.Username = result.Username
	p.Email = result.Email
	p.Password = result.Password
	p.Bio = result.Bio.String
	p.AvatarUrl = result.AvatarUrl.String
}
