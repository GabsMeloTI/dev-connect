-- name: CreateComment :one
INSERT INTO public."Comment"
(id, user_id, post_id, "content", likes, created_at)
VALUES(nextval('"Comment_id_seq"'::regclass), $1, $2, $3, 0, now())
    RETURNING *;

-- name: UpdateComment :one
UPDATE public."Comment"
SET "content"=$2
WHERE id=$1 AND user_id=$3
    RETURNING *;

-- name: IncrementCommentLikes :one
UPDATE public."Comment"
SET likes = likes + 1
WHERE id=$1
    RETURNING *;

-- name: DecrementCommentLikes :one
UPDATE public."Comment"
SET likes = likes - 1
WHERE id=$1 AND likes > 0
    RETURNING *;


-- name: DeleteComment :exec
DELETE FROM public."Comment"
WHERE id=$1 and user_id=$2 and post_id=$3;


-- name: GetAllComments :many
SELECT *
FROM "Comment"
WHERE post_id=$1
ORDER BY created_at DESC
    LIMIT $2 OFFSET $3;


-- name: GetCommentByID :one
SELECT *
FROM public."Comment"
WHERE id = $1;

