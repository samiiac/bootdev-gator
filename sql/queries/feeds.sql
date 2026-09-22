-- name: CreateFeed :one
INSERT INTO feeds(id,name,url,created_at,updated_at,user_id) 
VALUES($1,$2,$3,$4,$5,$6) RETURNING *;

-- name: GetAllFeed :many
SELECT * FROM feeds;

-- name: GetFeedByUrl :one
SELECT * from feeds WHERE url = $1;

-- name: MarkedFeedFetched :exec
UPDATE feeds SET last_fetched_at=$1,updated_at=$2 WHERE id=$3;

-- name: GetNextFeedToFetch :one
SELECT * from feeds ORDER BY last_fetched_at DESC LIMIT 1;