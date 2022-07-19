package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/iyorozuya/real-world-app/internal/sqlc"
	"github.com/iyorozuya/real-world-app/internal/types"
)

type FollowUserResponse struct {
	Profile types.Profile `json:"profile"`
}

func (s UserServiceImpl) Follow(params types.FollowUserParams) (*FollowUserResponse, error) {
	followUser, err := s.q.GetUserByName(context.Background(), params.Username)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("given profile %s not found", params.Username))
	}
	_ = s.q.FollowUser(context.Background(), sqlc.FollowUserParams{
		UserID:     int32(params.CurrentUser),
		FollowerID: followUser.ID,
	})
	return &FollowUserResponse{
		Profile: types.Profile{
			Username:  followUser.Username,
			Bio:       followUser.Bio.String,
			Image:     followUser.Image.String,
			Following: false,
		},
	}, nil
}
