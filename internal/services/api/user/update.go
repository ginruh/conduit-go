package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/iyorozuya/real-world-app/internal/services/api/auth"
	"github.com/iyorozuya/real-world-app/internal/sqlc"
	"github.com/iyorozuya/real-world-app/internal/types"
)

type UpdateUserResponse struct {
	User types.User `json:"user"`
}

func (s UserServiceImpl) Update(id int, params types.UpdateUserParams) (*UpdateUserResponse, error) {
	if params.Email == "" && params.Bio == "" && params.Image == "" {
		return nil, errors.New("email or bio or image field is required")
	}
	user, err := s.q.GetUserByID(context.Background(), int32(id))
	if err != nil {
		return nil, errors.New("unable to update user details")
	}
	email, bio, image := parseUserUpdateParams(user, params)
	updatedUser, err := s.q.UpdateUser(context.Background(), sqlc.UpdateUserParams{
		ID:    int32(id),
		Email: email,
		Bio:   bio,
		Image: image,
	})
	if err != nil {
		return nil, errors.New("unable to update user details")
	}
	token, err := auth.GenerateUserToken(int(user.ID))
	if err != nil {
		return nil, errors.New("internal server error")
	}
	return &UpdateUserResponse{
		User: types.User{
			Email:    updatedUser.Email,
			Token:    token,
			Username: updatedUser.Username,
			Bio:      updatedUser.Bio.String,
			Image:    updatedUser.Image.String,
		},
	}, nil
}

func parseUserUpdateParams(user sqlc.User, params types.UpdateUserParams) (string, sql.NullString, sql.NullString) {
	email := params.Email
	bio, image := sql.NullString{String: params.Bio, Valid: true}, sql.NullString{String: params.Image, Valid: true}
	if params.Email == "" {
		email = user.Email
	}
	if params.Bio == "" {
		bio = user.Bio
	}
	if params.Image == "" {
		image = user.Image
	}
	return email, bio, image
}
