package auth

import (
	"github.com/iyorozuya/real-world-app/internal/queries"
	"github.com/iyorozuya/real-world-app/internal/types"
)

type AuthService interface {
	Login(params types.LoginParams) (*LoginResponse, error)
	Register(params types.RegisterParams) (*RegisterUserResponse, error)
}

type AuthServiceImpl struct {
	q *queries.Queries
}

func NewAuthService(q *queries.Queries) AuthServiceImpl {
	return AuthServiceImpl{q}
}
