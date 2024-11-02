-- name: CreateComment :one
INSERT INTO public."Comment"
(id, user_id, post_id, "content", likes, created_at)
VALUES(nextval('"Comment_id_seq"'::regclass), $1, $2, $3, 0, now())
    RETURNING *;

-- name: UpdateComment :one
UPDATE public."Comment"
SET "content"=$2
WHERE id=$1
    RETURNING *;

-- name: UpdateCommentLikes :one
UPDATE public."Comment"
SET likes=$2
WHERE id=$1
    RETURNING *;

-- name: DeleteComment :exec
DELETE FROM public."Comment"
WHERE id=$1;

-- name: GetAllComments :many
SELECT *
FROM "Comment";