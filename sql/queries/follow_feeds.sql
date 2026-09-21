-- name: CreateFollowFeed :one
WITH inserted_follow_feed AS(
INSERT INTO follow_feeds(id,user_id,feed_id,created_at,updated_at)
VALUES($1,$2,$3,$4,$5) RETURNING *
)SELECT inserted_follow_feed.*,
feeds.name AS feed_name,
users.name AS user_name
from inserted_follow_feed INNER JOIN feeds ON inserted_follow_feed.feed_id = feeds.id
INNER JOIN users ON inserted_follow_feed.user_id = users.id;


-- name: GetFollowFeedsForUser :many
SELECT follow_feeds.*,feeds.name AS feed_name FROM follow_feeds 
INNER JOIN feeds ON follow_feeds.feed_id = feeds.id WHERE follow_feeds.user_id = $1;


-- name: UnfollowFeed :exec
DELETE from follow_feeds WHERE user_id =$1 and feed_id=$2;