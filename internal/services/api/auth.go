package api

import (
	"github.com/iyorozuya/real-world-app/internal/types"
	"github.com/iyorozuya/real-world-app/internal/utils"
)

type AuthService interface {
	Login() LoginResponse
	Register() RegisterUserResponse
}

type AuthServiceImpl struct {
	DB *utils.DB
}

type LoginResponse struct {
	User types.User `json:"user"`
}

func (service AuthServiceImpl) Login() LoginResponse {
	return LoginResponse{
		User: types.User{
			Email:    "iyaksha@tutanota.com",
			Token:    "abcdef",
			Username: "gintamashi",
			Bio:      "Silver soul",
			Image:    "",
		},
	}
}

type RegisterUserResponse struct {
	User types.User `json:"user"`
}

func (service AuthServiceImpl) Register() RegisterUserResponse {
	return RegisterUserResponse{
		User: types.User{
			Email:    "iyaksha@tutanota.com",
			Token:    "abcdef",
			Username: "gintamashi",
			Bio:      "Silver soul",
			Image:    "",
		},
	}
}
