package repository

import (
	"context"
	"database/sql"
	db "treads/db/sqlc"
)

type UserInterface interface {
	GetAllUsers(context.Context) ([]db.User, error)
	CreateUser(context.Context, db.CreateUserParams) (db.User, error)
	UpdateUser(context.Context, db.UpdateUserParams) (db.User, error)
	DeleteUser(context.Context, int64) error
	GetUsersLoginByEmail(context.Context, db.GetUsersLoginByEmailParams) (db.User, error)
	GetUsersLoginByUsername(context.Context, db.GetUsersLoginByUsernameParams) (db.User, error)
	GetUserById(context.Context, int64) (db.User, error)
	GetUsersByName(context.Context, string) (bool, error)
	GetUsersByEmail(context.Context, string) (bool, error)
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

func (r *User) GetAllUsers(ctx context.Context) ([]db.User, error) {
	return r.Queries.GetAllUsers(ctx)
}

func (r *User) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return r.Queries.CreateUser(ctx, arg)
}

func (r *User) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	return r.Queries.UpdateUser(ctx, arg)
}

func (r *User) DeleteUser(ctx context.Context, id int64) error {
	return r.Queries.DeleteUser(ctx, id)
}

func (r *User) GetUsersLoginByEmail(ctx context.Context, arg db.GetUsersLoginByEmailParams) (db.User, error) {
	return r.Queries.GetUsersLoginByEmail(ctx, arg)
}

func (r *User) GetUsersLoginByUsername(ctx context.Context, arg db.GetUsersLoginByUsernameParams) (db.User, error) {
	return r.Queries.GetUsersLoginByUsername(ctx, arg)
}

func (r *User) GetUserById(ctx context.Context, id int64) (db.User, error) {
	return r.Queries.GetUserById(ctx, id)
}

func (r *User) GetUsersByName(ctx context.Context, name string) (bool, error) {
	return r.Queries.GetUsersByUsername(ctx, name)
}

func (r *User) GetUsersByEmail(ctx context.Context, email string) (bool, error) {
	return r.Queries.GetUsersByEmail(ctx, email)
}
