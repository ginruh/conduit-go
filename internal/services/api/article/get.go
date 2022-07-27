package article

import (
	"errors"
	"fmt"
	"github.com/iyorozuya/real-world-app/internal/queries"
	"github.com/iyorozuya/real-world-app/internal/types"
	"strings"
)

type GetArticleResponse struct {
	Article types.Article
}

func (s ArticleServiceImpl) Get(params types.GetArticleParams) (*GetArticleResponse, error) {
	article, err := s.q.GetArticle(queries.GetArticleParams{
		Slug:          params.Slug,
		CurrentUserID: params.CurrentUser,
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
			CreatedAt:      article.CreatedAt.String(),
			UpdatedAt:      article.UpdatedAt.String(),
			Author: types.Author{
				Username:  article.AuthorUsername,
				Bio:       article.AuthorBio.String,
				Image:     article.AuthorImage.String,
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
