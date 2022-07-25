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
    CASE WHEN at.tags IS NULL THEN '' ELSE CAST(at.tags AS VARCHAR) END AS tags,
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
    LEFT JOIN
        (SELECT at.article_id, STRING_AGG(at.tag_name, ',') AS tags FROM article_tags AS at GROUP BY at.article_id) as at
    ON a.id = at.article_id AND at.article_id = a.id
WHERE a.slug = $1;

-- name: ListArticles :many
SELECT
    DISTINCT
    a.id,
    a.slug,
    a.title,
    a.description,
    a.body,
    u.username,
    u.bio,
    u.image,
    CASE WHEN af.favorites IS null THEN 0 ELSE CAST(af.favorites AS INTEGER) END AS favorites_count,
    CASE WHEN af_favorited.favorited IS null THEN FALSE ELSE TRUE END as favorited,
    CASE WHEN at.tags IS NULL THEN '' ELSE CAST(at.tags AS VARCHAR) END AS tags,
    a.created_at,
    a.updated_at
FROM article AS a
     LEFT JOIN "user" AS u ON a.author_id = u.id
     LEFT JOIN
        (
            SELECT uf.user_id, uf.follower_id, count(*) as following
            FROM user_follower as uf
            GROUP BY uf.user_id, uf.follower_id
        ) as uf
     ON uf.follower_id = a.author_id AND uf.user_id = sqlc.arg('CurrentUser')
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
     ON a.id = af_favorited.article_id AND af_favorited.user_id = sqlc.arg('CurrentUser')
     LEFT JOIN
     (
         SELECT af.article_id, af.user_id, u2.username
         FROM article_favorite as af
         LEFT JOIN "user" as u2 on u2.id = af.user_id
         GROUP BY af.article_id, af.user_id, u2.username
     ) as af_list
     ON
        a.id = af_list.article_id
            AND
        CASE
            WHEN sqlc.narg('FavoriteUsername') IS NULL THEN FALSE
            ELSE af_list.username IN(SELECT unnest(string_to_array(sqlc.narg('FavoriteUsername'), ',')) as parts)
        END
     LEFT JOIN
     (
         SELECT at.article_id, STRING_AGG(at.tag_name, ',') AS tags FROM article_tags AS at
         WHERE
             CASE
                 WHEN sqlc.narg('ArticleTag') IS NULL THEN TRUE
                 ELSE at.tag_name IN(SELECT unnest(string_to_array(sqlc.narg('ArticleTag'), ',')) as parts)
             END
         GROUP BY at.article_id
     ) AS at
    ON a.id = at.article_id AND at.article_id = a.id
WHERE
    CASE
        WHEN sqlc.narg('AuthorUsername') IS NULL THEN TRUE
        ELSE u.username IN(SELECT unnest(string_to_array(sqlc.narg('AuthorUsername'), ',')) as parts)
    END
  AND
    at.tags is not null
ORDER BY a.created_at DESC
LIMIT sqlc.arg('Limit')
OFFSET sqlc.arg('Offset');

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

