package utils

import (
	"errors"
	"github.com/iyorozuya/real-world-app/internal/services/api/auth"
	"strings"
)

func ParseAuthorizationHeader(tokenHeader string) (int, error) {
	tokenHeaderArr := strings.Split(tokenHeader, " ")
	if len(tokenHeaderArr) < 2 {
		return 0, errors.New("access forbidden")
	}
	if tokenHeaderArr[0] != "Token" {
		return 0, errors.New("access forbidden")
	}
	userId, err := auth.VerifyUserToken(tokenHeaderArr[1])
	if err != nil {
		return 0, errors.New("access forbidden")
	}
	return userId, nil
}
