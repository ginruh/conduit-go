-- name: GetUser :one
SELECT * FROM "user"
WHERE email = $1 LIMIT 1;

-- name: GetUserByName :one
SELECT * FROM "user"
WHERE username = $1 LIMIT 1;

-- name: GetUserByID :one
SELECT * FROM "user"
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO "user" (username, email, password)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateUser :one
UPDATE "user"
SET email = $2, bio = $3, image = $4
WHERE id = $1
RETURNING *;
