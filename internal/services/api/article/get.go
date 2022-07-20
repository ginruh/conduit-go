package article

import (
	"context"
	"errors"
	"fmt"
	"github.com/iyorozuya/real-world-app/internal/sqlc"
	"github.com/iyorozuya/real-world-app/internal/types"
	"strings"
)

type GetArticleResponse struct {
	Article types.Article
}

func (s ArticleServiceImpl) Get(params types.GetArticleParams) (*GetArticleResponse, error) {
	article, err := s.q.GetArticle(context.Background(), sqlc.GetArticleParams{
		Slug:   params.Slug,
		UserID: int32(params.CurrentUser),
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to find article %s", params.Slug))
	}
	return &GetArticleResponse{
		Article: types.Article{
			Slug:           article.Slug,
			Title:          article.Title,
			Description:    article.Description,
			Body:           article.Body,
			TagList:        parseArticleTags(article.Tags),
			Favorited:      article.Favorited,
			FavoritesCount: int(article.FavoritesCount),
			CreatedAt:      article.CreatedAt.Time.String(),
			UpdatedAt:      article.UpdatedAt.Time.String(),
			Author: types.Author{
				Username:  article.Username.String,
				Bio:       article.Bio.String,
				Image:     article.Image.String,
				Following: article.UserFollowing,
			},
		},
	}, nil
}

func parseArticleTags(tags string) []string {
	if tags == "" {
		return []string{}
	}
	return strings.Split(tags, ",")
}
