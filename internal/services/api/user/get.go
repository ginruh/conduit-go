package user

import "github.com/iyorozuya/real-world-app/internal/types"

type GetUserResponse struct {
	User types.User `json:"user"`
}

func (s UserServiceImpl) Get() {

}