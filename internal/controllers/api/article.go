package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/iyorozuya/real-world-app/internal/services/api/article"
	"github.com/iyorozuya/real-world-app/internal/types"
	"github.com/iyorozuya/real-world-app/internal/utils"
	"net/http"
)

type ArticleController struct {
	articleService article.ArticleService
	validate       *validator.Validate
}

func NewArticleController(articleService article.ArticleService, validate *validator.Validate) ArticleController {
	return ArticleController{
		articleService,
		validate,
	}
}

// List articles godoc
func (c ArticleController) List(w http.ResponseWriter, r *http.Request) {

}

// Feed (Feed for followed users) godoc
func (c ArticleController) Feed(w http.ResponseWriter, r *http.Request) {

}

// Get Article godoc
func (c ArticleController) Get(w http.ResponseWriter, r *http.Request) {
	getArticleParams := types.GetArticleParams{
		Slug:        chi.URLParam(r, "slug"),
		CurrentUser: r.Context().Value("userId").(int),
	}
	if validationErr := utils.ValidateStruct(c.validate, getArticleParams); validationErr != nil {
		utils.SendErrors(w, http.StatusUnprocessableEntity, validationErr)
		return
	}
	article, err := c.articleService.Get(getArticleParams)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendResponse(w, http.StatusOK, article)
}

// Create Article godoc
func (c ArticleController) Create(w http.ResponseWriter, r *http.Request) {

}

// Update Article godoc
func (c ArticleController) Update(w http.ResponseWriter, r *http.Request) {

}

// Delete Article godoc
func (c ArticleController) Delete(w http.ResponseWriter, r *http.Request) {

}

// Favorite Article godoc
func (c ArticleController) Favorite(w http.ResponseWriter, r *http.Request) {

}

// Unfavorite Article godoc
func (c ArticleController) Unfavorite(w http.ResponseWriter, r *http.Request) {

}
