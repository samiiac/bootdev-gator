-- +goose Up
CREATE TABLE follow_feeds(
    id UUID PRIMARY KEY,
    user_id  UUID NOT NULL,
    feed_id  UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT users_fk FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT feeds_fk FOREIGN KEY(feed_id) REFERENCES feeds(id) ON DELETE CASCADE,
    CONSTRAINT unique_records UNIQUE (user_id,feed_id)
);

-- +goose Down
DROP TABLE follow_feeds;