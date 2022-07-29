package queries

type FavoriteArticleParams struct {
	ArticleID string
	UserID    string
}

func (q Queries) FavoriteArticle(params FavoriteArticleParams) error {
	var err error
	_, err = q.db.NamedExec(
		`insert into article_favorite (article_id, user_id) VALUES (:article_id, :user_id)`,
		map[string]interface{}{
			"article_id": params.ArticleID,
			"user_id":    params.UserID,
		},
	)
	return err
}

type UnfavoriteArticleParams struct {
	ArticleID string
	UserID    string
}

func (q Queries) UnfavoriteArticle(params UnfavoriteArticleParams) error {
	var err error
	_, err = q.db.NamedExec(
		`delete from article_favorite where article_id = :article_id and user_id = :user_id`,
		map[string]interface{}{
			"article_id": params.ArticleID,
			"user_id":    params.UserID,
		},
	)
	return err
}
