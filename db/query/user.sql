-- name: CreateUser :one
INSERT INTO public."User" (id, "name", username, email, "password", bio, avatar_url, created_at, active)
VALUES(nextval('"User_id_seq"'::regclass), $1, $2, $3, $4, $5, $6, now(), true)
    RETURNING *;

-- name: UpdateUser :one
UPDATE "User"
SET "name"=$2, email=$3, username=$4, bio=$5, avatar_url=$6
WHERE id=$1
    RETURNING *;

-- name: UpdatePasswordByUserId :exec
UPDATE "User"
SET "password"=$2
WHERE id=$1;

-- name: DisableUser :exec
UPDATE "User"
SET active=false
WHERE id=$1;

-- name: DeleteUser :exec
DELETE FROM public."User"
WHERE id=$1;

-- name: GetUsersLoginByEmailOrUsername :one
SELECT *
FROM public."User"
WHERE active = true AND (email = $1 OR username = $1);


-- name: GetAllUsers :many
SELECT *
FROM "User"
WHERE active = true and username ILIKE '%' || $1 || '%'
order by username;


-- name: GetUserById :one
SELECT *
FROM public."User"
WHERE
    active = true AND
    id = $1;

-- name: GetUsersByUsernameOrEmail :one
SELECT EXISTS (
    SELECT 1
    FROM public."User"
    WHERE active = true
      AND (username = $1 OR email = $1)
);

