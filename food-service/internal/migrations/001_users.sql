-- +goose Up
CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE USERS;