package repository

import (
	"context"
	"database/sql"
	db "treads/db/sqlc"
)

type UserInterface interface {
	GetAllUsers(context.Context, sql.NullString) ([]db.User, error)
	CreateUser(context.Context, db.CreateUserParams) (db.User, error)
	UpdateUser(context.Context, db.UpdateUserParams) (db.User, error)
	UpdatePassword(ctx context.Context, arg db.UpdatePasswordByUserIdParams) error
	DisableUser(context.Context, int64) error
	DeleteUser(ctx context.Context, id int64) error
	GetUsersLoginByEmailOrUsername(context.Context, string) (db.User, error)
	GetUserById(context.Context, int64) (db.User, error)
	GetUsersByUsernameOrEmail(context.Context, string) (bool, error)
}

type User struct {
	DBtx    db.DBTX
	Queries *db.Queries
	SqlConn *sql.DB
}

func NewUser(sqlDB *sql.DB) *User {
	q := db.New(sqlDB)
	return &User{
		DBtx:    sqlDB,
		Queries: q,
		SqlConn: sqlDB,
	}
}

func (r *User) GetAllUsers(ctx context.Context, username sql.NullString) ([]db.User, error) {
	return r.Queries.GetAllUsers(ctx, username)
}

func (r *User) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return r.Queries.CreateUser(ctx, arg)
}

func (r *User) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	return r.Queries.UpdateUser(ctx, arg)
}

func (r *User) UpdatePassword(ctx context.Context, arg db.UpdatePasswordByUserIdParams) error {
	return r.Queries.UpdatePasswordByUserId(ctx, arg)
}

func (r *User) DisableUser(ctx context.Context, id int64) error {
	return r.Queries.DeleteUser(ctx, id)
}

func (r *User) DeleteUser(ctx context.Context, id int64) error {
	return r.Queries.DeleteUser(ctx, id)
}

func (r *User) GetUsersLoginByEmailOrUsername(ctx context.Context, arg string) (db.User, error) {
	return r.Queries.GetUsersLoginByEmailOrUsername(ctx, arg)
}

func (r *User) GetUserById(ctx context.Context, id int64) (db.User, error) {
	return r.Queries.GetUserById(ctx, id)
}

func (r *User) GetUsersByUsernameOrEmail(ctx context.Context, arg string) (bool, error) {
	return r.Queries.GetUsersByUsernameOrEmail(ctx, arg)
}
