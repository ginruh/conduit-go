package article

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/iyorozuya/real-world-app/internal/queries"
	"github.com/iyorozuya/real-world-app/internal/types"
	"strings"
)

type CreateArticleResponse struct {
	Article types.Article
}

func (s ArticleServiceImpl) Create(params types.CreateArticleParams) (*CreateArticleResponse, error) {
	article, err := s.q.CreateArticle(queries.CreateArticleParams{
		Slug:        generateArticleSlug(params.Article.Title),
		Title:       params.Article.Title,
		Description: params.Article.Description,
		Body:        params.Article.Body,
		AuthorID:    params.CurrentUser,
	})
	if err != nil {
		return nil, errors.New("unable to save article")
	}
	if len(params.Article.TagList) > 0 {
		err = s.q.CreateArticleTags(queries.CreateArticleTagsParams{
			ArticleID: article.ID,
			Tags:      params.Article.TagList,
		})
		if err != nil {
			return nil, errors.New("unable to save article tags")
		}
	}
	article, _ = s.q.GetArticle(queries.GetArticleParams{
		Slug: article.Slug,
		CurrentUserID: sql.NullString{
			String: params.CurrentUser,
			Valid:  true,
		},
	})
	return &CreateArticleResponse{
		Article: types.Article{
			Slug:           article.Slug,
			Title:          article.Title,
			Description:    article.Description,
			Body:           article.Body,
			TagList:        parseArticleTags(article.Tags),
			Favorited:      article.Favorited,
			FavoritesCount: article.FavoritesCount,
			CreatedAt:      article.CreatedAt.String(),
			UpdatedAt:      article.UpdatedAt.String(),
			Author: types.Author{
				Username:  article.AuthorUsername,
				Bio:       article.AuthorBio.String,
				Image:     article.AuthorImage.String,
				Following: false,
			},
		},
	}, nil
}

func generateArticleSlug(title string) string {
	var slug string
	titleSplitVal := strings.Split(strings.ToLower(title), " ")
	for i, word := range titleSplitVal {
		slug = fmt.Sprintf("%s%s", slug, word)
		if i < len(titleSplitVal)-1 {
			slug = fmt.Sprintf("%s-", slug)
		}
	}
	return slug
}
