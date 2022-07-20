-- +goose Up
-- +goose StatementBegin
CREATE TABLE "article_comment" (
    id SERIAL PRIMARY KEY,
    article_id INT NOT NULL REFERENCES "article" (id),
    user_id INT NOT NULL REFERENCES "user" (id),
    body TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
DROP TABLE "article_comment";
