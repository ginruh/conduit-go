package user

import (
	"github.com/iyorozuya/real-world-app/internal/queries"
	"github.com/iyorozuya/real-world-app/internal/types"
)

type UserService interface {
	Get(id string) (*GetUserResponse, error)
	Update(id string, params types.UpdateUserParams) (*UpdateUserResponse, error)
	GetProfile(params types.GetProfileParams) (*GetProfileResponse, error)
	Follow(params types.FollowUserParams) (*FollowUserResponse, error)
	Unfollow(params types.UnfollowUserParams) (*UnfollowUserResponse, error)
}

type UserServiceImpl struct {
	q *queries.Queries
}

func NewUserService(q *queries.Queries) UserServiceImpl {
	return UserServiceImpl{q}
}
