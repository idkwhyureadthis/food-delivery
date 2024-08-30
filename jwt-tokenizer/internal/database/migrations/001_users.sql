-- +goose Up
CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY,
    hashed_password TEXT NOT NULL,
    name TEXT NOT NULL,
    access_token TEXT NOT NULL
);

-- +goose Down
DROP TABLE USERS;