package user

import (
	"github.com/iyorozuya/real-world-app/internal/sqlc"
	"github.com/iyorozuya/real-world-app/internal/types"
)

type UserService interface {
	Get(id int) (*GetUserResponse, error)
	Update(id int, params types.UpdateUserParams) (*UpdateUserResponse, error)
	GetProfile(params types.GetProfileParams) (*GetProfileResponse, error)
}

type UserServiceImpl struct {
	q *sqlc.Queries
}

func NewUserService(q *sqlc.Queries) UserServiceImpl {
	return UserServiceImpl{q}
}
