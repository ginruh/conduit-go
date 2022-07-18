package user

import "github.com/iyorozuya/real-world-app/internal/sqlc"

type UserService interface {
	Get(id string) (*GetUserResponse, error)
	Update()
}

type UserServiceImpl struct {
	q *sqlc.Queries
}

func NewUserService(q *sqlc.Queries) UserServiceImpl {
	return UserServiceImpl{q}
}
