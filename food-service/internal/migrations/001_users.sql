-- +goose Up
CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY,
    hashed_password TEXT NOT NULL,
    name TEXT NOT NULL,
    refresh_token TEXT
);

-- +goose Down
DROP TABLE USERS;