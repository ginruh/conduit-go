package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/iyorozuya/real-world-app/internal/queries"
	"github.com/iyorozuya/real-world-app/internal/types"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type LoginResponse struct {
	User types.User `json:"user"`
}

func (service AuthServiceImpl) Login(params types.LoginParams) (*LoginResponse, error) {
	user, err := service.q.GetUserByEmail(queries.GetUserByEmailParams{
		Email: params.Email,
	})
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	token, err := GenerateUserToken(user.ID)
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

func GenerateUserToken(userID string) (string, error) {
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

func VerifyUserToken(tokenString string) (string, error) {
	type tokenClaims struct {
		ID string `json:"id"`
		jwt.RegisteredClaims
	}
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		jwtSecret := os.Getenv("JWT_SECRET")
		return []byte(jwtSecret), nil
	})
	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		return claims.ID, nil
	}
	return "", err
}
