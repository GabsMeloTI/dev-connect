-- name: CreateUser :one
INSERT INTO public."User" (id, "name", username, email, "password", bio, avatar_url, created_at, active)
VALUES(nextval('"User_id_seq"'::regclass), $1, $2, $3, $4, $5, $6, now(), true)
    RETURNING *;

-- name: UpdateUser :one
UPDATE "User"
SET "name"=$2, email=$3, username=$4, "password"=$5, bio=$6, avatar_url=$7
WHERE id=$1
    RETURNING *;

-- name: DeleteUser :exec
UPDATE "User"
SET active=false
WHERE id=$1;

-- name: GetUsersLoginByEmail :one
SELECT *
FROM public."User"
WHERE active = true and email = $1 and "password" = $2;

-- name: GetUsersLoginByUsername :one
SELECT *
FROM public."User"
WHERE active = true and username = $1 and "password" = $2;

-- name: GetAllUsers :many
SELECT *
FROM "User"
WHERE active = true;

-- name: GetUserById :one
SELECT *
FROM public."User"
WHERE
    active = true AND
    id = $1;

-- name: GetUsersByUsername :one
SELECT EXISTS(
    SELECT *
    FROM public."User"
    WHERE
        active = true and
        username = $1
);

-- name: GetUsersByEmail :one
SELECT EXISTS(
    SELECT *
    FROM public."User"
    WHERE
        active = true and
        "name" = $1
);