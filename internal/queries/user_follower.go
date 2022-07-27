package queries

import "github.com/iyorozuya/real-world-app/internal/models"

type GetFollowersParams struct {
	UserID string
}

func (q Queries) GetFollowers(params GetFollowersParams) ([]models.UserFollower, error) {
	var userFollowers []models.UserFollower
	rows, err := q.db.Queryx(
		`SELECT * FROM user_follower WHERE user_id = ?`,
		params.UserID,
	)
	if err != nil {
		return []models.UserFollower{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var userFollower models.UserFollower
		err = rows.Scan(&userFollower)
		if err != nil {
			break
		}
		userFollowers = append(userFollowers, userFollower)
	}
	if err != nil {
		return []models.UserFollower{}, err
	}
	return userFollowers, err
}

type FollowUserParams struct {
	UserID     string
	FollowerID string
}

func (q Queries) FollowUser(params FollowUserParams) error {
	_, err := q.db.Exec(
		"INSERT INTO user_follower (user_id, follower_id) VALUES (?, ?)",
		params.UserID,
		params.FollowerID,
	)
	if err != nil {
		return err
	}
	return nil
}

type UnfollowUserParams struct {
	UserID     string
	FollowerID string
}

func (q Queries) UnfollowUser(params UnfollowUserParams) error {
	_, err := q.db.Exec(
		"DELETE FROM user_follower WHERE user_id = ? AND follower_id = ?",
		params.UserID,
		params.FollowerID,
	)
	if err != nil {
		return err
	}
	return nil
}
