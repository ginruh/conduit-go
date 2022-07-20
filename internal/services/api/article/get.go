package article

import "github.com/iyorozuya/real-world-app/internal/types"

type GetArticleResponse struct {
	Article types.Article
}

func (s ArticleServiceImpl) Get(params types.GetArticleParams) (*GetArticleResponse, error) {
	return nil, nil
}
