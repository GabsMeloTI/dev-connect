// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO public."User" (id, "name", username, email, "password", bio, avatar_url, created_at, active)
VALUES(nextval('"User_id_seq"'::regclass), $1, $2, $3, $4, $5, $6, now(), true)
    RETURNING id, name, username, email, password, bio, avatar_url, active, created_at, last_login
`

type CreateUserParams struct {
	Name      string         `json:"name"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Bio       sql.NullString `json:"bio"`
	AvatarUrl sql.NullString `json:"avatar_url"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Name,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.Bio,
		arg.AvatarUrl,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Bio,
		&i.AvatarUrl,
		&i.Active,
		&i.CreatedAt,
		&i.LastLogin,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
UPDATE "User"
SET active=false
WHERE id=$1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, name, username, email, password, bio, avatar_url, active, created_at, last_login
FROM "User"
WHERE active = true
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.Bio,
			&i.AvatarUrl,
			&i.Active,
			&i.CreatedAt,
			&i.LastLogin,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserById = `-- name: GetUserById :one
SELECT id, name, username, email, password, bio, avatar_url, active, created_at, last_login
FROM public."User"
WHERE
    active = true AND
    id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Bio,
		&i.AvatarUrl,
		&i.Active,
		&i.CreatedAt,
		&i.LastLogin,
	)
	return i, err
}

const getUsersByEmail = `-- name: GetUsersByEmail :one
SELECT EXISTS(
    SELECT id, name, username, email, password, bio, avatar_url, active, created_at, last_login
    FROM public."User"
    WHERE
        active = true and
        "name" = $1
)
`

func (q *Queries) GetUsersByEmail(ctx context.Context, name string) (bool, error) {
	row := q.db.QueryRowContext(ctx, getUsersByEmail, name)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getUsersByUsername = `-- name: GetUsersByUsername :one
SELECT EXISTS(
    SELECT id, name, username, email, password, bio, avatar_url, active, created_at, last_login
    FROM public."User"
    WHERE
        active = true and
        username = $1
)
`

func (q *Queries) GetUsersByUsername(ctx context.Context, username string) (bool, error) {
	row := q.db.QueryRowContext(ctx, getUsersByUsername, username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getUsersLoginByEmail = `-- name: GetUsersLoginByEmail :one
SELECT id, name, username, email, password, bio, avatar_url, active, created_at, last_login
FROM public."User"
WHERE active = true and email = $1 and "password" = $2
`

type GetUsersLoginByEmailParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) GetUsersLoginByEmail(ctx context.Context, arg GetUsersLoginByEmailParams) (User, error) {
	row := q.db.QueryRowContext(ctx, getUsersLoginByEmail, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Bio,
		&i.AvatarUrl,
		&i.Active,
		&i.CreatedAt,
		&i.LastLogin,
	)
	return i, err
}

const getUsersLoginByUsername = `-- name: GetUsersLoginByUsername :one
SELECT id, name, username, email, password, bio, avatar_url, active, created_at, last_login
FROM public."User"
WHERE active = true and username = $1 and "password" = $2
`

type GetUsersLoginByUsernameParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (q *Queries) GetUsersLoginByUsername(ctx context.Context, arg GetUsersLoginByUsernameParams) (User, error) {
	row := q.db.QueryRowContext(ctx, getUsersLoginByUsername, arg.Username, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Bio,
		&i.AvatarUrl,
		&i.Active,
		&i.CreatedAt,
		&i.LastLogin,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE "User"
SET "name"=$2, email=$3, username=$4, "password"=$5, bio=$6, avatar_url=$7
WHERE id=$1
    RETURNING id, name, username, email, password, bio, avatar_url, active, created_at, last_login
`

type UpdateUserParams struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	Bio       sql.NullString `json:"bio"`
	AvatarUrl sql.NullString `json:"avatar_url"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Username,
		arg.Password,
		arg.Bio,
		arg.AvatarUrl,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Bio,
		&i.AvatarUrl,
		&i.Active,
		&i.CreatedAt,
		&i.LastLogin,
	)
	return i, err
}
