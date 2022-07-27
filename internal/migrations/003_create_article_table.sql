-- +goose Up
-- +goose StatementBegin
CREATE TABLE article (
    id VARCHAR(36) PRIMARY KEY,
    slug VARCHAR(255) NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    body TEXT NOT NULL,
    author_id VARCHAR(36) NOT NULL REFERENCES user (id),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
DROP TABLE article;