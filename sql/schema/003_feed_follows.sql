-- +goose Up
CREATE TABLE feed_follows (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE feed_id INTEGER NOT NULL REFERENCES feeds (id) ON DELETE CASCADE ADD CONSTRAINT user_feed_pair UNIQUE (userid, feed_id)
);

-- +goose Down
Drop Table feed_follows
