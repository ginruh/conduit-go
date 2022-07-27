-- +goose Up
-- +goose StatementBegin
CREATE TABLE article_comment (
    id INT PRIMARY KEY AUTO_INCREMENT,
    article_id INT NOT NULL REFERENCES article (id),
    user_id INT NOT NULL REFERENCES user (id),
    body TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
DROP TABLE article_comment;
