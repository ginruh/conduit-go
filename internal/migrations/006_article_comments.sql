-- +goose Up
-- +goose StatementBegin
CREATE TABLE article_comment (
    id VARCHAR(36) PRIMARY KEY,
    article_id VARCHAR(36) NOT NULL REFERENCES article (id),
    user_id VARCHAR(36) NOT NULL REFERENCES user (id),
    body TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
DROP TABLE article_comment;
