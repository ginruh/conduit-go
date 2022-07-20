-- +goose Up
-- +goose StatementBegin
CREATE TABLE "article_tags" (
    article_id INT NOT NULL REFERENCES article (id),
    tag_name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (article_id, tag_name)
);
-- +goose StatementEnd

-- +goose.Down
DROP TABLE article_tags;