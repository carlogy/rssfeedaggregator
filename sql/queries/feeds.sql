-- name: CreateFeed :one
INSERT INTO feeds(created_at, updated_at, name, url, user_id)
VALUES (
$1,
$2,
$3,
$4,
$5
)
RETURNING *;

-- name: GetFeeds :many
SELECT
    *
FROM
    feeds;


-- name: GetFeedByURL :one

SELECT
    *
FROM
    feeds
WHERE
    feeds.url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET
    updated_at = $1,
    last_fetched_at = $2
WHERE
    feeds.id = $3;


-- name: GetNextFeedToFetch :one
-- WITH oldest_fetched AS (
-- SELECT
--     f.last_fetched_at
-- FROM
--     feeds f
-- ORDER BY f.last_fetched_at ASC NULLS LAST
-- Limit 1
-- )
-- SELECT
--     *
-- From
--     feeds f
-- WHERE
--     f.last_fetched_at IS NULL OR
--     f.last_fetched_at <= (SELECT last_fetched_at FROM oldest_fetched)
-- ORDER BY f.last_fetched_at ASC NULLS FIRST
-- Limit 1;

SELECT
    *
FROM
    feeds f
ORDER BY f.last_fetched_at ASC NULLS FIRST
LIMIT 1;
