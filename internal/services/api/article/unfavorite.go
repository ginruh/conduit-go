package article

import (
	"database/sql"
	"errors"
	"github.com/iyorozuya/real-world-app/internal/queries"
	"github.com/iyorozuya/real-world-app/internal/types"
)

type UnfavoriteArticleResponse struct {
	Article types.Article
}

func (s ArticleServiceImpl) Unfavorite(params types.UnfavoriteArticleParams) (*UnfavoriteArticleResponse, error) {
	article, err := s.q.GetArticle(queries.GetArticleParams{
		Slug: params.Slug,
		CurrentUserID: sql.NullString{
			String: params.CurrentUser,
			Valid:  true,
		},
	})
	if err != nil {
		return nil, errors.New("article not found")
	}
	err = s.q.UnfavoriteArticle(queries.UnfavoriteArticleParams{
		ArticleID: article.ID,
		UserID:    params.CurrentUser,
	})
	if err != nil {
		return nil, errors.New("unable to favorite article")
	}
	return &UnfavoriteArticleResponse{
		Article: types.Article{
			Slug:           article.Slug,
			Title:          article.Title,
			Description:    article.Description,
			Body:           article.Body,
			TagList:        parseArticleTags(article.Tags),
			Favorited:      false,
			FavoritesCount: article.FavoritesCount - 1,
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
