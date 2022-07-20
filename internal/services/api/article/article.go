package article

import "github.com/iyorozuya/real-world-app/internal/sqlc"

type ArticleService interface {
	List()
	Feed()
	Get(id string)
	Create()
	Update()
	Delete()
	Favorite()
	Unfavorite()
	ListTags()
}

type ArticleServiceImpl struct {
	q *sqlc.Queries
}

func NewArticleService(q *sqlc.Queries) ArticleServiceImpl {
	return ArticleServiceImpl{q}
}
