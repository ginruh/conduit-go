package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string         `db:"id"`
	Email     string         `db:"email"`
	Password  string         `db:"password"`
	Username  string         `db:"username"`
	Bio       sql.NullString `db:"bio"`
	Image     sql.NullString `db:"image"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}
