package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/iyorozuya/real-world-app/internal/services/api/comment"
	"net/http"
)

type CommentController struct {
	commentService comment.CommentService
	validate       *validator.Validate
}

func NewCommentController(commentService comment.CommentService, validate *validator.Validate) CommentController {
	return CommentController{
		commentService,
		validate,
	}
}

// List Comments godoc
func (c CommentController) List(w http.ResponseWriter, r *http.Request) {

}

// Create Comment godoc
func (c CommentController) Create(w http.ResponseWriter, r *http.Request) {

}

// Delete Comment godoc
func (c CommentController) Delete(w http.ResponseWriter, r *http.Request) {

}
