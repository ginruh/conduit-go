-- +goose Up
-- +goose StatementBegin
CREATE TABLE article_tags (
    article_id BINARY(16) NOT NULL REFERENCES article (id),
    tag_name VARCHAR(50) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (article_id, tag_name)
);
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS article_tags;
