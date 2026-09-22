-- name: CreatePost :one
INSERT INTO posts(id,title,created_at,updated_at,url,description,published_at,feed_id)
VALUES($1,$2,$3,$4,$5,$6,$7,$8) RETURNING *;


-- name: GetPostForUser :many
SELECT * from posts JOIN follow_feeds ON follow_feeds.feed_id = posts.feed_id WHERE follow_feeds.user_id = $1
ORDER BY posts.published_at DESC LIMIT $2;