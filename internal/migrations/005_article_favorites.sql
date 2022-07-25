-- +goose Up
-- +goose StatementBegin
CREATE TABLE article_favorite (
    article_id INT NOT NULL REFERENCES article (id),
    user_id INT NOT NULL REFERENCES user (id),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (article_id, user_id)
);
-- +goose StatementEnd

-- +goose Down
DROP TABLE "article_favorite";
