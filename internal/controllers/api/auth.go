package api

import (
	"github.com/iyorozuya/real-world-app/internal/services/api"
	"github.com/iyorozuya/real-world-app/internal/utils"
	"net/http"
)

type AuthController struct {
	AuthService api.AuthService
}

// Register godoc
// @Summary  User registration
// @Tags     Users
// @Produce  application/json
// @Router   /users [post]
func (c AuthController) Register(w http.ResponseWriter, r *http.Request) {
	registerUserResponse := c.AuthService.Register()
	utils.SendResponse(w, http.StatusOK, registerUserResponse)
}

// Login godoc
// @Summary  User login
// @Tags     Users
// @Produce  application/json
// @Router   /users/login [post]
func (c AuthController) Login(w http.ResponseWriter, r *http.Request) {
	loginResponse := c.AuthService.Login()
	utils.SendResponse(w, http.StatusOK, loginResponse)
}
