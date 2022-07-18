package auth

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/iyorozuya/real-world-app/internal/types"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type LoginResponse struct {
	User types.User `json:"user"`
}

func (service AuthServiceImpl) Login(params types.LoginParams) (*LoginResponse, error) {
	user, err := service.q.GetUser(context.Background(), params.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	token, err := generateUserToken(user.Password, params.Password, int(user.ID))
	if err != nil {
		return nil, err
	}
	return &LoginResponse{
		User: types.User{
			Email:    params.Email,
			Token:    token,
			Username: user.Username,
			Bio:      user.Bio.String,
			Image:    user.Image.String,
		},
	}, nil
}

func generateUserToken(hashedPassword, password string, userID int) (string, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userID,
	})
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
