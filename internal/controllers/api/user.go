package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/iyorozuya/real-world-app/internal/services/api/user"
	"net/http"
)

type UserController struct {
	userService user.UserService
	validate    *validator.Validate
}

func NewUserController(userService user.UserService, validate *validator.Validate) UserController {
	return UserController{userService, validate}
}

// GetCurrentUser godoc
func (c UserController) GetCurrentUser(w http.ResponseWriter, r *http.Request) {

}

// UpdateUser godoc
func (c UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// GetProfile godoc
func (c UserController) GetProfile(w http.ResponseWriter, r *http.Request) {

}

// FollowUser godoc
func (c UserController) FollowUser(w http.ResponseWriter, r *http.Request) {

}

// UnfollowUser godoc
func (c UserController) UnfollowUser(w http.ResponseWriter, r *http.Request) {

}
