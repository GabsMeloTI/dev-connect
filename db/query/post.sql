-- name: CreatePost :one
INSERT INTO public."Post"
(id, user_id, "content", image_url, likes, shares, created_at, archive)
VALUES(nextval('"Post_id_seq"'::regclass), $1, $2, $3, 0, 0, now(), false)
    RETURNING *;

-- name: UpdatePost :one
UPDATE public."Post"
SET "content"=$2, image_url=$3
WHERE id=$1
    RETURNING *;

-- name: ArchivePost :exec
UPDATE "Post"
SET archive=true
WHERE id=$1;

-- name: DeletePost :exec
DELETE FROM public."Post"
WHERE id=$1;

-- name: GetAllPosts :many
SELECT p.id,
       p."content",
       p.image_url,
       p.likes,
       p.shares,
       p.created_at,
       p.archive,
       p.user_id,
       u.name,
       u.avatar_url
FROM "Post" p
         INNER JOIN "User" u ON u.id = p.user_id
WHERE p.archive = true;


-- name: GetAllPostsByUser :many
SELECT
    p.id,
    p."content",
    p.image_url,
    p.likes,
    p.shares,
    p.created_at,
    p.archive,
    p.user_id,
    u.name,
    u.avatar_url
FROM "Post" p
    INNER JOIN "User" u ON u.id = p.user_id
WHERE user_id=$1;