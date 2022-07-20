package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/iyorozuya/real-world-app/internal/services/api/article"
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
	c.articleService.Get()
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
