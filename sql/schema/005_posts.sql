-- +goose Up
CREATE table posts (
    id UUID primary key,
    title TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    url  TEXT UNIQUE NOT NULL,
    description varchar(100),
    published_at TIMESTAMP,
    feed_id UUID NOT NULL,
    CONSTRAINT feed_fk FOREIGN KEY(feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;