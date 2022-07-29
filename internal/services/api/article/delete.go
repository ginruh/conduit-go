package article

import (
	"errors"
	"github.com/iyorozuya/real-world-app/internal/queries"
	"github.com/iyorozuya/real-world-app/internal/types"
)

type DeleteArticleResponse struct {
	Slug string
}

func (s ArticleServiceImpl) Delete(params types.DeleteArticleParams) (*DeleteArticleResponse, error) {
	articleSlug, err := s.q.DeleteArticle(queries.DeleteArticleParams{
		Slug: params.Slug,
	})
	if err != nil {
		return nil, errors.New("unable to delete article")
	}
	return &DeleteArticleResponse{
		Slug: articleSlug,
	}, nil
}
