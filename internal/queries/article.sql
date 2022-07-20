-- name: GetArticle :one
SELECT
    a.slug,
    a.title,
    a.description,
    a.body,
    u.username,
    u.bio,
    u.image,
    CASE WHEN uf.following IS null THEN FALSE ELSE TRUE END AS user_following,
    CASE WHEN af.favorites IS null THEN 0 ELSE CAST(af.favorites AS INTEGER) END AS favorites_count,
    CASE WHEN af_favorited.favorited IS null THEN FALSE ELSE TRUE END as favorited,
    a.created_at,
    a.updated_at
FROM article AS a
    LEFT JOIN "user" AS u ON a.author_id = u.id
    LEFT JOIN
        (SELECT uf.user_id, uf.follower_id, count(*) as following FROM user_follower as uf GROUP BY uf.user_id, uf.follower_id) as uf
    ON uf.follower_id = a.author_id AND uf.user_id = $2
    LEFT JOIN
        (SELECT article_id, COUNT(*) as favorites FROM article_favorite as af GROUP BY af.article_id) as af
    ON a.id = af.article_id
    LEFT JOIN
        (SELECT af.article_id, af.user_id, COUNT(*) as favorited FROM article_favorite as af GROUP BY af.article_id, af.user_id) as af_favorited
    ON a.id = af_favorited.article_id AND af_favorited.user_id = $2
WHERE a.slug = $1;

-- name: CreateArticle :one
INSERT INTO "article" (slug, title, description, body, author_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateArticle :one
UPDATE "article"
SET slug = $2, title = $3, description = $4, body = $5
WHERE id = $1
RETURNING *;

-- name: DeleteArticle :one
DELETE FROM "article"
WHERE slug = $1
RETURNING *;

