package user

import (
	"context"
	"errors"
	"github.com/iyorozuya/real-world-app/internal/services/api/auth"
	"github.com/iyorozuya/real-world-app/internal/types"
)

type GetUserResponse struct {
	User types.User `json:"user"`
}

func (s UserServiceImpl) Get(id int) (*GetUserResponse, error) {
	user, err := s.q.GetUserByID(context.Background(), int32(id))
	if err != nil {
		return nil, errors.New("user not found")
	}
	token, err := auth.GenerateUserToken(id)
	if err != nil {
		return nil, errors.New("internal server error")
	}
	return &GetUserResponse{
		User: types.User{
			Email:    user.Email,
			Token:    token,
			Username: user.Username,
			Bio:      user.Bio.String,
			Image:    user.Image.String,
		},
	}, nil
}
