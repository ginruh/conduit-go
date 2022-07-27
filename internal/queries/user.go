package queries

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/iyorozuya/real-world-app/internal/models"
)

type GetUserByIdParams struct {
	ID string
}

func (q Queries) GetUserById(params GetUserByIdParams) (models.User, error) {
	var user models.User
	row := q.db.QueryRowx(
		`SELECT * FROM user WHERE id = ?`,
		params.ID,
	)
	if err := row.StructScan(&user); err != nil {
		return user, err
	}
	return user, nil
}

type GetUserByEmailParams struct {
	Email string
}

func (q Queries) GetUserByEmail(params GetUserByEmailParams) (models.User, error) {
	var user models.User
	row := q.db.QueryRowx(
		`SELECT * FROM user WHERE email = ?`,
		params.Email,
	)
	if err := row.StructScan(&user); err != nil {
		return user, err
	}
	return user, nil
}

type GetUserByNameParams struct {
	Username string
}

func (q Queries) GetUserByName(params GetUserByNameParams) (models.User, error) {
	var user models.User
	row := q.db.QueryRowx(
		`SELECT * FROM user WHERE username = ?`,
		params.Username,
	)
	if err := row.StructScan(&user); err != nil {
		return user, err
	}
	return user, nil
}

type CreateUserParams struct {
	Username string
	Email    string
	Password string
}

func (q Queries) CreateUser(params CreateUserParams) (models.User, error) {
	userId := uuid.New()
	_, err := q.db.Exec(
		`INSERT INTO user (id, username, email, password) VALUES (?, ?, ?, ?)`,
		userId.String(),
		params.Username,
		params.Email,
		params.Password,
	)
	if err != nil {
		return models.User{}, err
	}
	user, _ := q.GetUserById(GetUserByIdParams{ID: userId.String()})
	return user, nil
}

type UpdateUserParams struct {
	ID    string
	Email string
	Bio   sql.NullString
	Image sql.NullString
}

func (q Queries) UpdateUser(params UpdateUserParams) (models.User, error) {
	_, err := q.db.Exec(
		`UPDATE user SET email = ?, bio = ?, image = ? WHERE id = ?`,
		params.Email,
		params.Bio,
		params.Image,
		params.ID,
	)
	if err != nil {
		return models.User{}, err
	}
	user, _ := q.GetUserByEmail(GetUserByEmailParams{Email: params.Email})
	return user, nil
}
