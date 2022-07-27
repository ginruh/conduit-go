package article

import (
	"errors"
	"github.com/iyorozuya/real-world-app/internal/queries"
	"github.com/iyorozuya/real-world-app/internal/types"
	"strconv"
	"strings"
)

type ListArticlesResponse struct {
	Articles     []types.Article `json:"articles"`
	ArticleCount int             `json:"articleCount"`
}

func (s ArticleServiceImpl) List(params types.ListArticlesParams) (*ListArticlesResponse, error) {
	var articlesList []types.Article
	articlesLimit, _ := strconv.Atoi(params.Limit.String)
	if articlesLimit == 0 {
		// Default article count
		articlesLimit = 20
	}
	articlesOffset, _ := strconv.Atoi(params.Offset.String)
	articles, err := s.q.ListArticles(queries.ListArticlesParams{
		CurrentUser:    params.CurrentUser,
		FavoritedUsers: params.Favorited,
		ArticleTags:    params.Tag,
		AuthorUsername: params.Author,
		Limit:          articlesLimit,
		Offset:         articlesOffset,
	})
	if err != nil {
		return nil, errors.New("unable to list articles")
	}
	for _, article := range articles {
		articlesList = append(articlesList, types.Article{
			Slug:           article.Slug,
			Title:          article.Title,
			Description:    article.Description,
			Body:           article.Body,
			TagList:        strings.Split(article.Tags, ","),
			Favorited:      article.Favorited,
			FavoritesCount: article.FavoritesCount,
			CreatedAt:      article.CreatedAt.String(),
			UpdatedAt:      article.UpdatedAt.String(),
			Author: types.Author{
				Username:  article.AuthorUsername,
				Bio:       article.AuthorBio.String,
				Image:     article.AuthorImage.String,
				Following: article.UserFollowing,
			},
		})
	}
	return &ListArticlesResponse{
		Articles:     articlesList,
		ArticleCount: len(articlesList),
	}, nil
}
