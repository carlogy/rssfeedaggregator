-- name: CreatePost :one
INSERT INTO posts(created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
$1,
$2,
$3,
$4,
$5,
$6,
$7
)
RETURNING *;

-- name: GetPosts :many
SELECT
    *
FROM
    posts p
ORDER BY
    p.published_at DESC NULLS LAST
LIMIT $1;
