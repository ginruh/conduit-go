package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/iyorozuya/real-world-app/internal/sqlc"
	"github.com/iyorozuya/real-world-app/internal/types"
)

type UnfollowUserResponse struct {
	Profile types.Profile `json:"profile"`
}

func (s UserServiceImpl) Unfollow(params types.UnfollowUserParams) (*UnfollowUserResponse, error) {
	followUser, err := s.q.GetUserByName(context.Background(), params.Username)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("given profile %s not found", params.Username))
	}
	err = s.q.UnfollowUser(context.Background(), sqlc.UnfollowUserParams{
		UserID:     int32(params.CurrentUser),
		FollowerID: followUser.ID,
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to unfollow %s", params.Username))
	}
	return &UnfollowUserResponse{
		Profile: types.Profile{
			Username:  followUser.Username,
			Bio:       followUser.Bio.String,
			Image:     followUser.Image.String,
			Following: false,
		},
	}, nil
}
