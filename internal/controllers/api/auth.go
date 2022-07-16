package api

import (
	"github.com/iyorozuya/real-world-app/internal/services/api"
	"net/http"
)

type AuthController struct {
	AuthService api.AuthService
}

// Register godoc
// @Summary  User registration
// @Tags     users
// @Produce  text/html
// @Router   /api/users/ [post]
func (c AuthController) Register(w http.ResponseWriter, r *http.Request) {
	c.AuthService.Register()
}

// Login godoc
// @Summary  User login
// @Tags     users
// @Produce  text/html
// @Router   /api/users/login [post]
func (c AuthController) Login(w http.ResponseWriter, r *http.Request) {
	c.AuthService.Login()
}
