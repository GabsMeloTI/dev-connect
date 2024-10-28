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
DELETE FROM public."User"
WHERE id=$1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const disableUser = `-- name: DisableUser :exec
UPDATE "User"
SET active=false
WHERE id=$1
`

func (q *Queries) DisableUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, disableUser, id)
	return err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, name, username, email, password, bio, avatar_url, active, created_at, last_login
FROM "User"
WHERE active = true and username ILIKE '%' || $1 || '%'
order by username
`

func (q *Queries) GetAllUsers(ctx context.Context, dollar_1 sql.NullString) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers, dollar_1)
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

const getUsersByUsernameOrEmail = `-- name: GetUsersByUsernameOrEmail :one
SELECT EXISTS (
    SELECT 1
    FROM public."User"
    WHERE active = true
      AND (username = $1 OR email = $1)
)
`

func (q *Queries) GetUsersByUsernameOrEmail(ctx context.Context, username string) (bool, error) {
	row := q.db.QueryRowContext(ctx, getUsersByUsernameOrEmail, username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getUsersLoginByEmailOrUsername = `-- name: GetUsersLoginByEmailOrUsername :one
SELECT id, name, username, email, password, bio, avatar_url, active, created_at, last_login
FROM public."User"
WHERE active = true and email = $1 or username = $1
`

func (q *Queries) GetUsersLoginByEmailOrUsername(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUsersLoginByEmailOrUsername, email)
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

const updatePasswordByUserId = `-- name: UpdatePasswordByUserId :exec
UPDATE "User"
SET "password"=$2
WHERE id=$1
`

type UpdatePasswordByUserIdParams struct {
	ID       int64  `json:"id"`
	Password string `json:"password"`
}

func (q *Queries) UpdatePasswordByUserId(ctx context.Context, arg UpdatePasswordByUserIdParams) error {
	_, err := q.db.ExecContext(ctx, updatePasswordByUserId, arg.ID, arg.Password)
	return err
}

const updateUser = `-- name: UpdateUser :one
UPDATE "User"
SET "name"=$2, email=$3, username=$4, bio=$5, avatar_url=$6
WHERE id=$1
    RETURNING id, name, username, email, password, bio, avatar_url, active, created_at, last_login
`

type UpdateUserParams struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Username  string         `json:"username"`
	Bio       sql.NullString `json:"bio"`
	AvatarUrl sql.NullString `json:"avatar_url"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Username,
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
