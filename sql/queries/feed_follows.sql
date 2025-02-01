-- Add a CreateFeedFollow query. It will be a deceptively complex SQL query. It should insert a feed follow record, but then return all the fields from the field follow as well as the names of the linked user and feed. I'll add a tip at the bottom of this lesson if you need it.
-- name CreateFeedFollow :one
INSERT INTO
    feed_follows ()
