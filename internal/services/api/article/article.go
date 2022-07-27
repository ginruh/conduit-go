package article

import (
	"github.com/iyorozuya/real-world-app/internal/queries"
	"github.com/iyorozuya/real-world-app/internal/types"
)

type ArticleService interface {
	Get(params types.GetArticleParams) (*GetArticleResponse, error)
	List(params types.ListArticlesParams) (*ListArticlesResponse, error)
	// Feed()
	Create(params types.CreateArticleParams) (*CreateArticleResponse, error)
	// Update()
	// Delete()
	// Favorite()
	// Unfavorite()
	// ListTags()
}

type ArticleServiceImpl struct {
	q *queries.Queries
}

func NewArticleService(q *queries.Queries) ArticleServiceImpl {
	return ArticleServiceImpl{q}
}
