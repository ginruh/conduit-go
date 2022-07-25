package article

import "github.com/iyorozuya/real-world-app/internal/types"

type ListArticlesResponse struct {
	Articles     []types.Article `json:"articles"`
	ArticleCount int             `json:"articleCount"`
}

func (s ArticleServiceImpl) List(params types.ListArticlesParams) (*ListArticlesResponse, error) {

}
