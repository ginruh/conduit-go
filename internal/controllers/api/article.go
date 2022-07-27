package api

import (
	"database/sql"
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
	listArticleParams := types.ListArticlesParams{
		Tag:         r.URL.Query().Get("tag"),
		Author:      r.URL.Query().Get("author"),
		Favorited:   r.URL.Query().Get("favorited"),
		Limit:       r.URL.Query().Get("limit"),
		Offset:      r.URL.Query().Get("offset"),
		CurrentUser: r.Context().Value("userId").(sql.NullString),
	}
	if validationErr := utils.ValidateStruct(c.validate, listArticleParams); validationErr != nil {
		utils.SendErrors(w, http.StatusUnprocessableEntity, validationErr)
		return
	}
	articles, err := c.articleService.List(listArticleParams)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendResponse(w, http.StatusOK, articles)
}

// Feed (Feed for followed users) godoc
func (c ArticleController) Feed(w http.ResponseWriter, r *http.Request) {

}

// Get Article godoc
func (c ArticleController) Get(w http.ResponseWriter, r *http.Request) {
	getArticleParams := types.GetArticleParams{
		Slug:        chi.URLParam(r, "slug"),
		CurrentUser: r.Context().Value("userId").(sql.NullString),
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
	var createArticleParams types.CreateArticleParams
	if validationErr := utils.ValidateBody(r.Body, c.validate, &createArticleParams); validationErr != nil {
		utils.SendErrors(w, http.StatusUnprocessableEntity, validationErr)
		return
	}
	userId := r.Context().Value("userId").(string)
	createArticleParams.CurrentUser = userId
	article, err := c.articleService.Create(createArticleParams)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendResponse(w, http.StatusOK, article)
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
