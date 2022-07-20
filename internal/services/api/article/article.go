package article

import (
	"github.com/iyorozuya/real-world-app/internal/sqlc"
	"github.com/iyorozuya/real-world-app/internal/types"
)

type ArticleService interface {
	Get(params types.GetArticleParams) (GetArticleResponse, error)
	// List()
	// Feed()
	// Create()
	// Update()
	// Delete()
	// Favorite()
	// Unfavorite()
	// ListTags()
}

type ArticleServiceImpl struct {
	q *sqlc.Queries
}

func NewArticleService(q *sqlc.Queries) ArticleServiceImpl {
	return ArticleServiceImpl{q}
}
