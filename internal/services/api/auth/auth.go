package auth

import (
	"github.com/iyorozuya/real-world-app/internal/sqlc"
	"github.com/iyorozuya/real-world-app/internal/types"
)

type AuthService interface {
	Login(params types.LoginParams) (*LoginResponse, error)
	Register(params types.RegisterParams) (*RegisterUserResponse, error)
}

type AuthServiceImpl struct {
	q *sqlc.Queries
}

func NewAuthService(q *sqlc.Queries) AuthServiceImpl {
	return AuthServiceImpl{q}
}
