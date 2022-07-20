-- +goose Up
-- +goose StatementBegin
CREATE TABLE "article" (
    id SERIAL PRIMARY KEY,
    slug VARCHAR(255) NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE "article"
ADD author_id INT NOT NULL REFERENCES "user" (id);
-- +goose StatementEnd

-- +goose Down
DROP TABLE "article";