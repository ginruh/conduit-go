package article

import (
	"database/sql"
	"errors"
	"github.com/iyorozuya/real-world-app/internal/queries"
	"github.com/iyorozuya/real-world-app/internal/types"
)

type FavoriteArticleResponse struct {
	Article types.Article
}

func (s ArticleServiceImpl) Favorite(params types.FavoriteArticleParams) (*FavoriteArticleResponse, error) {
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
	err = s.q.FavoriteArticle(queries.FavoriteArticleParams{
		ArticleID: article.ID,
		UserID:    params.CurrentUser,
	})
	if err != nil {
		return nil, errors.New("unable to favorite article")
	}
	return &FavoriteArticleResponse{
		Article: types.Article{
			Slug:           article.Slug,
			Title:          article.Title,
			Description:    article.Description,
			Body:           article.Body,
			TagList:        parseArticleTags(article.Tags),
			Favorited:      true,
			FavoritesCount: article.FavoritesCount + 1,
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
