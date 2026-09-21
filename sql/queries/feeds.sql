-- name: CreateFeed :one
INSERT INTO feeds(id,name,url,created_at,updated_at,user_id) 
VALUES($1,$2,$3,$4,$5,$6) RETURNING *;

-- name: GetAllFeed :many
SELECT * FROM feeds;

-- name: GetFeedByUrl :one
SELECT * from feeds WHERE url = $1;