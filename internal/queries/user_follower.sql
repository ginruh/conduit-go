-- name: GetFollowers :many
SELECT * FROM user_follower
WHERE user_id = $1;

-- name: FollowUser :exec
INSERT INTO user_follower (user_id, follower_id)
VALUES ($1, $2);

-- name: UnfollowUser :exec
DELETE FROM user_follower
WHERE user_id = $1 AND follower_id = $2;
