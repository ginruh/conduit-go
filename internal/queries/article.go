package queries

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/iyorozuya/real-world-app/internal/models"
)

type GetArticleParams struct {
	Slug          string
	CurrentUserID sql.NullString
}

func (q Queries) GetArticle(params GetArticleParams) (models.ArticleDetails, error) {
	var articleDetails models.ArticleDetails
	var err error
	rows, err := q.db.NamedQuery(
		`
		SELECT
		    a.id,
		    a.slug,
		    a.title,
		    a.description,
		    a.body,
		    u.username,
		    u.bio,
		    u.image,
		    IF(uf.following IS null, FALSE, TRUE) AS user_following,
		    IF(af.favorites IS null, 0, CONVERT(af.favorites, UNSIGNED)) AS favorites_count,
		    IF(af_favorited.favorited IS null, FALSE, TRUE) as favorited,
		    IF(at.tags IS NULL, NULL, CONVERT(at.tags, CHAR)) AS tags,
		    a.created_at,
		    a.updated_at
		FROM article AS a
		    LEFT JOIN user AS u ON a.author_id = u.id
		    LEFT JOIN
		        (SELECT uf.user_id, uf.follower_id, count(*) as following FROM user_follower as uf GROUP BY uf.user_id, uf.follower_id) as uf
		    ON uf.follower_id = a.author_id AND uf.user_id = :current_user_id
		    LEFT JOIN
		        (SELECT article_id, COUNT(*) as favorites FROM article_favorite as af GROUP BY af.article_id) as af
		    ON a.id = af.article_id
		    LEFT JOIN
		        (SELECT af.article_id, af.user_id, COUNT(*) as favorited FROM article_favorite as af GROUP BY af.article_id, af.user_id) as af_favorited
		    ON a.id = af_favorited.article_id AND af_favorited.user_id = :current_user_id
		    LEFT JOIN
		        (SELECT at.article_id, GROUP_CONCAT(DISTINCT at.tag_name SEPARATOR ',') AS tags FROM article_tags AS at GROUP BY at.article_id) as at
		    ON a.id = at.article_id AND at.article_id = a.id
		WHERE a.slug = :article_slug
		LIMIT 1
		`,
		map[string]interface{}{
			"current_user_id": params.CurrentUserID,
			"article_slug":    params.Slug,
		},
	)
	if err != nil {
		return models.ArticleDetails{}, err
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.StructScan(&articleDetails); err != nil {
			return models.ArticleDetails{}, nil
		}
	}
	return articleDetails, nil
}

type ListArticlesParams struct {
	CurrentUser    sql.NullString
	FavoritedUsers sql.NullString
	ArticleTags    sql.NullString
	AuthorUsername sql.NullString
	Limit          int
	Offset         int
}

func (q Queries) ListArticles(params ListArticlesParams) ([]models.ArticleDetails, error) {
	var articlesList []models.ArticleDetails
	var err error
	rows, err := q.db.NamedQuery(
		`
			SELECT
			    a.id,
			    a.slug,
			    a.title,
			    a.description,
			    a.body,
			    u.username,
			    u.bio,
			    u.image,
		    	IF(uf.following IS null, FALSE, TRUE) AS user_following,
			    IF(af.favorites IS null, 0, CONVERT(af.favorites, UNSIGNED)) AS favorites_count,
			    IF(af_favorited.favorited IS null, FALSE, TRUE) as favorited,
			    IF(at.tags IS NULL, NULL, CONVERT(at.tags, CHAR)) AS tags,
			    a.created_at,
			    a.updated_at
			FROM article AS a
			     LEFT JOIN user AS u ON a.author_id = u.id
			     LEFT JOIN
			        (
			            SELECT uf.user_id, uf.follower_id, count(*) as following
			            FROM user_follower as uf
			            GROUP BY uf.user_id, uf.follower_id
			        ) as uf
			     ON uf.follower_id = a.author_id AND uf.user_id = :current_user_id
			     LEFT JOIN
			        (
			            SELECT article_id, COUNT(*) as favorites
			            FROM article_favorite as af
			            GROUP BY af.article_id
			        ) as af
			     ON a.id = af.article_id
			     LEFT JOIN
			     (
			         SELECT af.article_id, af.user_id, COUNT(*) as favorited
			         FROM article_favorite as af
			         GROUP BY af.article_id, af.user_id
			     ) as af_favorited
			     ON a.id = af_favorited.article_id AND af_favorited.user_id = :current_user_id
			     LEFT JOIN
			     (
			         SELECT af.article_id, af.user_id, u2.username
			         FROM article_favorite as af
			         LEFT JOIN user as u2 on u2.id = af.user_id
			         GROUP BY af.article_id, af.user_id, u2.username
			     ) as af_list
			     ON
			        a.id = af_list.article_id AND
			        IF (:favorited_users is NULL, FALSE, FIND_IN_SET(af_list.username, :favorited_users))
			     LEFT JOIN
			     (
			         SELECT at.article_id, GROUP_CONCAT(DISTINCT at.tag_name SEPARATOR ',') AS tags FROM article_tags AS at
			         WHERE IF (:article_tags IS NULL, TRUE, FIND_IN_SET(at.tag_name, :article_tags))
			         GROUP BY at.article_id
			     ) AS at
			    ON a.id = at.article_id AND at.article_id = a.id
			WHERE IF (:author_usernames IS NULL, TRUE, FIND_IN_SET(u.username, :author_usernames))
			ORDER BY a.created_at DESC
			LIMIT :limit
			OFFSET :offset;
		`,
		map[string]interface{}{
			"current_user_id":  params.CurrentUser,
			"favorited_users":  params.FavoritedUsers,
			"article_tags":     params.ArticleTags,
			"author_usernames": params.AuthorUsername,
			"limit":            params.Limit,
			"offset":           params.Offset,
		},
	)
	if err != nil {
		return articlesList, err
	}
	defer rows.Close()
	for rows.Next() {
		var articleDetails models.ArticleDetails
		if err = rows.StructScan(&articleDetails); err != nil {
			break
		}
		articlesList = append(articlesList, articleDetails)
	}
	if err != nil {
		return []models.ArticleDetails{}, err
	}
	return articlesList, nil
}

type CreateArticleParams struct {
	Slug        string
	Title       string
	Description string
	Body        string
	AuthorID    string
}

func (q Queries) CreateArticle(params CreateArticleParams) (models.ArticleDetails, error) {
	articleId := uuid.New()
	_, err := q.db.NamedExec(
		`
			INSERT INTO article (id, slug, title, description, body, author_id)
			VALUES (:article_id, :article_slug, :article_title, :article_description, :article_body, :author_id)
		`,
		map[string]interface{}{
			"article_id":          articleId,
			"article_slug":        params.Slug,
			"article_title":       params.Title,
			"article_description": params.Description,
			"article_body":        params.Body,
			"author_id":           params.AuthorID,
		},
	)
	if err != nil {
		return models.ArticleDetails{}, err
	}
	article, err := q.GetArticle(GetArticleParams{
		Slug: params.Slug,
		CurrentUserID: sql.NullString{
			String: params.AuthorID,
			Valid:  true,
		},
	})
	if err != nil {
		return models.ArticleDetails{}, err
	}
	return article, nil
}

type UpdateArticleParams struct {
	ID          string
	Slug        string
	Title       string
	Description string
	Body        string
	AuthorID    string
}

func (q Queries) UpdateArticle(params UpdateArticleParams) (models.ArticleDetails, error) {
	_, err := q.db.NamedExec(
		`
			UPDATE article SET 
			    slug = :article_slug, 
			    title = :article_title, 
			    description = :article_description,
			    body = :article_body
			WHERE id = :article_id;
		`,
		map[string]interface{}{
			"article_id":          params.ID,
			"article_slug":        params.Slug,
			"article_title":       params.Title,
			"article_description": params.Description,
			"article_body":        params.Body,
		},
	)
	if err != nil {
		return models.ArticleDetails{}, err
	}
	article, err := q.GetArticle(GetArticleParams{
		Slug: params.Slug,
		CurrentUserID: sql.NullString{
			String: params.AuthorID,
			Valid:  true,
		},
	})
	if err != nil {
		return models.ArticleDetails{}, err
	}
	return article, nil
}

type DeleteArticleParams struct {
	Slug string
}

func (q Queries) DeleteArticle(params DeleteArticleParams) (string, error) {
	_, err := q.db.NamedExec(
		`DELETE FROM article WHERE slug = :article_slug`,
		map[string]interface{}{
			"article_slug": params.Slug,
		},
	)
	if err != nil {
		return "", err
	}
	return params.Slug, nil
}
