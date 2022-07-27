package utils

import (
	"errors"
	"github.com/iyorozuya/real-world-app/internal/services/api/auth"
	"strings"
)

func ParseAuthorizationHeader(tokenHeader string) (string, error) {
	tokenHeaderArr := strings.Split(tokenHeader, " ")
	if len(tokenHeaderArr) < 2 {
		return "", errors.New("access forbidden")
	}
	if tokenHeaderArr[0] != "Token" {
		return "", errors.New("access forbidden")
	}
	userId, err := auth.VerifyUserToken(tokenHeaderArr[1])
	if err != nil {
		return "", errors.New("access forbidden")
	}
	return userId, nil
}
