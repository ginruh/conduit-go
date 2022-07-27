package types

import "database/sql"

type RegisterParams struct {
	Username string `json:"username" type:"string" validate:"required"`
	Email    string `json:"email" type:"string" validate:"required,email"`
	Password string `json:"password" type:"string" validate:"required"`
}

type LoginParams struct {
	Email    string `json:"email" type:"string" validate:"required,email"`
	Password string `json:"password" type:"string" validate:"required"`
}

type UpdateUserParams struct {
	Email string `json:"email" type:"string" validate:"omitempty,email"`
	Bio   string `json:"bio" type:"string" validate:"omitempty,min=8,max=60"`
	Image string `json:"image" type:"string" validate:"omitempty,url"`
}

type GetProfileParams struct {
	Username          string         `json:"username" type:"string" validate:"required,alphanum"`
	AuthenticatedUser sql.NullString `validate:"omitempty"`
}

type FollowUserParams struct {
	Username    string `json:"username" type:"string" validate:"required,alphanum"`
	CurrentUser string `validate:"omitempty"`
}

type UnfollowUserParams struct {
	Username    string `json:"username" type:"string" validate:"required,alphanum"`
	CurrentUser string `validate:"omitempty"`
}

type GetArticleParams struct {
	Slug        string         `json:"slug" type:"string" validate:"required,lowercase"`
	CurrentUser sql.NullString `validate:"omitempty"`
}

type ListArticlesParams struct {
	Tag         sql.NullString `type:"string" validate:"omitempty,alpha,lowercase"`
	Author      sql.NullString `type:"string" validate:"omitempty,alphanum"`
	Favorited   sql.NullString `type:"string" validate:"omitempty,alphanum"`
	Limit       sql.NullString `validate:"omitempty,number"`
	Offset      sql.NullString `validate:"omitempty,number"`
	CurrentUser sql.NullString `validate:"omitempty"`
}

type CreateArticleParams struct {
	Article     CreateArticleStructParams `json:"article"`
	CurrentUser string                    `validate:"omitempty"`
}

type CreateArticleStructParams struct {
	Title       string   `json:"title" validate:"required,min=8,max=50"`
	Description string   `json:"description" validate:"required,min=8,max=255"`
	Body        string   `json:"body" validate:"required,min=8"`
	TagList     []string `json:"tagList" validate:"omitempty"`
}
