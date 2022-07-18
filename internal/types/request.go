package types

type RegisterParams struct {
	Username string `json:"username" type:"string" validate:"required"`
	Email    string `json:"email" type:"string" validate:"required,email"`
	Password string `json:"password" type:"string" validate:"required"`
}

type LoginParams struct {
	Email    string `json:"email" type:"string" validate:"required,email"`
	Password string `json:"password" type:"string" validate:"required"`
}