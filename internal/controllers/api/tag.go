package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/iyorozuya/real-world-app/internal/services/api/tag"
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

}
