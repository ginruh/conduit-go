package models

import "time"

type UserFollower struct {
	UserID     string    `db:"user_id"`
	FollowerID string    `db:"follower_id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
