-- name: GetUser :one
SELECT * FROM "user"
WHERE email = $1 OR username = $1 LIMIT 1;

-- name: GetUserByID :one
SELECT * FROM "user"
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO "user" (username, email, password)
VALUES ($1, $2, $3)
RETURNING *;