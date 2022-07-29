package queries

type CreateArticleTagsParams struct {
	ArticleID string
	Tags      []string
}

func (q Queries) CreateArticleTags(params CreateArticleTagsParams) error {
	var err error
	for _, tagName := range params.Tags {
		_, err = q.db.NamedExec(
			`insert into article_tags (article_id, tag_name) values (:article_id, :article_tag_name)`,
			map[string]interface{}{
				"article_id":       params.ArticleID,
				"article_tag_name": tagName,
			},
		)
		if err != nil {
			break
		}
	}
	return err
}

func (q Queries) ListArticleTags() ([]string, error) {
	var articleTags []string
	rows, err := q.db.Queryx(`select distinct tag_name from article_tags`)
	for rows.Next() {
		var articleTag string
		if err = rows.Scan(&articleTag); err != nil {
			break
		}
		articleTags = append(articleTags, articleTag)
	}
	if err != nil {
		return []string{}, err
	}
	return articleTags, nil
}
