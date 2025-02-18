-- +goose Up
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    description TEXT,
    published_at TIMESTAMP NOT NULL,
    feed_id INTEGER NOT NULL REFERENCES feeds (id) ON DELETE CASCADE
);

-- +goose Down
Drop Table posts;
