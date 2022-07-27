package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/iyorozuya/real-world-app/internal/services/api/user"
	"github.com/iyorozuya/real-world-app/internal/types"
	"github.com/iyorozuya/real-world-app/internal/utils"
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
	userId := r.Context().Value("userId").(string)
	user, err := c.userService.Get(userId)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendResponse(w, http.StatusOK, user)
}

// UpdateUser godoc
func (c UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUserParams types.UpdateUserParams
	if validationErr := utils.ValidateBody(r.Body, c.validate, &updateUserParams); validationErr != nil {
		utils.SendErrors(w, http.StatusUnprocessableEntity, validationErr)
		return
	}
	userId := r.Context().Value("userId").(string)
	user, err := c.userService.Update(userId, updateUserParams)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendResponse(w, http.StatusOK, user)
}

// GetProfile godoc
func (c UserController) GetProfile(w http.ResponseWriter, r *http.Request) {
	getProfileParams := types.GetProfileParams{
		Username:          chi.URLParam(r, "username"),
		AuthenticatedUser: r.Context().Value("userId").(string),
	}
	if validationErr := utils.ValidateStruct(c.validate, &getProfileParams); validationErr != nil {
		utils.SendErrors(w, http.StatusUnprocessableEntity, validationErr)
		return
	}
	profile, err := c.userService.GetProfile(getProfileParams)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendResponse(w, http.StatusOK, profile)
}

// FollowUser godoc
func (c UserController) FollowUser(w http.ResponseWriter, r *http.Request) {
	followUserParams := types.FollowUserParams{
		Username:    chi.URLParam(r, "username"),
		CurrentUser: r.Context().Value("userId").(string),
	}
	if validationErr := utils.ValidateStruct(c.validate, &followUserParams); validationErr != nil {
		utils.SendErrors(w, http.StatusUnprocessableEntity, validationErr)
		return
	}
	profile, err := c.userService.Follow(followUserParams)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendResponse(w, http.StatusOK, profile)
}

// UnfollowUser godoc
func (c UserController) UnfollowUser(w http.ResponseWriter, r *http.Request) {
	unfollowUserParams := types.UnfollowUserParams{
		Username:    chi.URLParam(r, "username"),
		CurrentUser: r.Context().Value("userId").(string),
	}
	if validationErr := utils.ValidateStruct(c.validate, &unfollowUserParams); validationErr != nil {
		utils.SendErrors(w, http.StatusUnprocessableEntity, validationErr)
		return
	}
	profile, err := c.userService.Unfollow(unfollowUserParams)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendResponse(w, http.StatusOK, profile)
}
