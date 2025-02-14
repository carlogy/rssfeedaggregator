-- name: CreateFeedFollow :one
WITH inserted_feed_follows AS (
INSERT INTO
    feed_follows(created_at, updated_at, user_id, feed_id)
    VALUES (
    $1,
    $2,
    $3,
    $4
    )
    RETURNING *
    )

    SELECT
        inserted_feed_follows.*,
        feeds.name AS feed_name,
        users.name AS user_name
    FROM
        inserted_feed_follows
    INNER JOIN
        feeds ON inserted_feed_follows.feed_id = feeds.id
    INNER JOIN
        users ON inserted_feed_follows.user_id = users.id;


-- name: GetFeedFollowsForUser :many
WITH get_feed_follows AS (
SELECT
    *
FROM
    feed_follows
WHERE
    feed_follows.user_id = $1
)

SELECT
    get_feed_follows.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM
    get_feed_follows
INNER JOIN feeds on get_feed_follows.feed_id = feeds.id
INNER JOIN users on get_feed_follows.user_id = users.id;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
USING feeds
WHERE
    feeds.id = feed_follows.feed_id
AND
    feed_follows.user_id = $1
AND
    feeds.url = $2;
