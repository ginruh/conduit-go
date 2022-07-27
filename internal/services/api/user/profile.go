package user

import (
	"errors"
	"fmt"
	"github.com/iyorozuya/real-world-app/internal/queries"
	"github.com/iyorozuya/real-world-app/internal/types"
)

type GetProfileResponse struct {
	Profile types.Profile `json:"profile"`
}

func (s UserServiceImpl) GetProfile(params types.GetProfileParams) (*GetProfileResponse, error) {
	username := params.Username
	user, err := s.q.GetUserByName(queries.GetUserByNameParams{
		Username: username,
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf(`cannot get profile %s`, username))
	}
	return &GetProfileResponse{
		Profile: types.Profile{
			Username:  user.Username,
			Bio:       user.Bio.String,
			Image:     user.Image.String,
			Following: false,
		},
	}, nil
}
