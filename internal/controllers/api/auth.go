package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/iyorozuya/real-world-app/internal/services/api/auth"
	"github.com/iyorozuya/real-world-app/internal/types"
	"github.com/iyorozuya/real-world-app/internal/utils"
	"net/http"
)

type AuthController struct {
	authService auth.AuthService
	validate    *validator.Validate
}

func NewAuthController(authService auth.AuthService, validate *validator.Validate) AuthController {
	return AuthController{
		authService, validate,
	}
}

// Register godoc
// @Summary  User registration
// @Tags     Users
// @Param    body  body  RegisterParams  true  "register user params"
// @Accept   json
// @Produce  json
// @Success  200  {object}  auth.RegisterUserResponse
// @Failure  422  {object}  utils.ErrorResponse
// @Failure  500  {object}  utils.ErrorResponse
// @Router   /users [post]
func (c AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var registerParams types.RegisterParams
	if validationErr := utils.ValidateBody(r.Body, c.validate, &registerParams); validationErr != nil {
		utils.SendErrors(w, http.StatusUnprocessableEntity, validationErr)
		return
	}
	registerUserResponse, err := c.authService.Register(registerParams)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendResponse(w, http.StatusOK, *registerUserResponse)
}

// Login godoc
// @Summary  User login
// @Tags     Users
// @Param    body  body  LoginParams  true  "login user params"
// @Accept   json
// @Produce  json
// @Success  200  {object}  auth.LoginResponse
// @Failure  422  {object}  utils.ErrorResponse
// @Failure  403  {object}  utils.ErrorResponse
// @Router   /users/login [post]
func (c AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var loginParams types.LoginParams
	if validationErr := utils.ValidateBody(r.Body, c.validate, &loginParams); validationErr != nil {
		utils.SendErrors(w, http.StatusUnprocessableEntity, validationErr)
		return
	}
	loginResponse, err := c.authService.Login(loginParams)
	if err != nil {
		utils.SendError(w, http.StatusForbidden, err.Error())
		return
	}
	utils.SendResponse(w, http.StatusOK, *loginResponse)
}
