package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/iyorozuya/real-world-app/internal/services/api/tag"
	"github.com/iyorozuya/real-world-app/internal/utils"
	"net/http"
)

type TagController struct {
	tagService tag.TagService
	validate   *validator.Validate
}

func NewTagController(tagService tag.TagService, validate *validator.Validate) TagController {
	return TagController{
		tagService,
		validate,
	}
}

// List Tags godoc
func (c TagController) List(w http.ResponseWriter, r *http.Request) {
	tags, err := c.tagService.List()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendResponse(w, http.StatusOK, tags)
}
