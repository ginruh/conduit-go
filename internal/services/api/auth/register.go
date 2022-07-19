package auth

import (
	"context"
	"errors"
	"github.com/iyorozuya/real-world-app/internal/sqlc"
	"github.com/iyorozuya/real-world-app/internal/types"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
)

type RegisterUserResponse struct {
	User types.User `json:"user"`
}

func (service AuthServiceImpl) Register(params types.RegisterParams) (*RegisterUserResponse, error) {
	bcryptCost, _ := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	user, err := service.q.CreateUser(context.Background(), sqlc.CreateUserParams{
		Email:    params.Email,
		Username: params.Username,
		Password: string(passwordHash),
	})
	if err != nil {
		return nil, errors.New("user already exists")
	}
	token, err := GenerateUserToken(int(user.ID))
	if err != nil {
		return nil, errors.New("internal server error")
	}
	return &RegisterUserResponse{
		User: types.User{
			Email:    user.Email,
			Token:    token,
			Username: user.Username,
			Bio:      user.Bio.String,
			Image:    user.Image.String,
		},
	}, nil
}
