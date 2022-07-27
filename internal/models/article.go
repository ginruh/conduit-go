package models

import (
	"database/sql"
	"time"
)

type ArticleDetails struct {
	ID             string         `db:"id"`
	Slug           string         `db:"slug"`
	Title          string         `db:"title"`
	Description    string         `db:"description"`
	Body           string         `db:"body"`
	AuthorUsername string         `db:"username"`
	AuthorBio      sql.NullString `db:"bio"`
	AuthorImage    sql.NullString `db:"image"`
	UserFollowing  bool           `db:"user_following"`
	FavoritesCount int            `db:"favorites_count"`
	Favorited      bool           `db:"favorited"`
	Tags           string         `db:"tags"`
	CreatedAt      time.Time      `db:"created_at"`
	UpdatedAt      time.Time      `db:"updated_at"`
}

type ListArticleDetails struct {
	ID             string         `db:"id"`
	Slug           string         `db:"slug"`
	Title          string         `db:"title"`
	Description    string         `db:"description"`
	Body           string         `db:"body"`
	AuthorUsername string         `db:"username"`
	AuthorBio      sql.NullString `db:"bio"`
	AuthorImage    sql.NullString `db:"image"`
	FavoritesCount int            `db:"favorites_count"`
	Favorited      bool           `db:"favorited"`
	Tags           string         `db:"tags"`
	CreatedAt      time.Time      `db:"created_at"`
	UpdatedAt      time.Time      `db:"updated_at"`
}
