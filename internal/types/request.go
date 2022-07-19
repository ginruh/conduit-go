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

type UpdateUserParams struct {
	Email string `json:"email" type:"string" validate:"omitempty,email"`
	Bio   string `json:"bio" type:"string" validate:"omitempty,min=8,max=60"`
	Image string `json:"image" type:"string" validate:"omitempty,url"`
}

type GetProfileParams struct {
	Username          string `json:"username" type:"string" validate:"required,alphanum"`
	AuthenticatedUser int    `validate:"omitempty"`
}
