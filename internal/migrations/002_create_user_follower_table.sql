-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user_follower" (
    user_id INT NOT NULL REFERENCES "user" (id),
    follower_id INT NOT NULL REFERENCES "user" (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, follower_id)
);
-- +goose StatementEnd

-- +goose Down
DROP TABLE "user_follower";
