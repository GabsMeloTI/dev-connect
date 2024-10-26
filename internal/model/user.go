package model

import (
	"database/sql"
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
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Bio       string `json:"bio"`
	AvatarUrl string `json:"avatar_url"`
}

type UserUpdateDto struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Bio       string `json:"bio"`
	AvatarUrl string `json:"avatar_url"`
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
		Password: p.Password,
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

func (p *UserResponse) ParseFromUserObject(result db.User) {
	p.ID = result.ID
	p.Name = result.Name
	p.Username = result.Username
	p.Email = result.Email
	p.Password = result.Password
	p.Bio = result.Bio.String
	p.AvatarUrl = result.AvatarUrl.String
}
