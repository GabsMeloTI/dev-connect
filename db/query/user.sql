-- name: CreateUser :one
INSERT INTO public."User" (id, "name", username, email, "password", bio, avatar_url, created_at, last_login, active)
VALUES(nextval('"User_id_seq"'::regclass), $1, $2, $3, $4, $5, $6, now(), $2, true)
    RETURNING *;

-- name: UpdateUser :one
UPDATE "User"
SET "name"=$2, email=$3, "password"=$4, bio=$5, avatar_url=$6
WHERE id=$1
    RETURNING *;

-- name: DeleteUser :exec
UPDATE "User"
SET active=false
WHERE id=$1;

-- name: GetAllUsers :many
SELECT *
FROM "User";