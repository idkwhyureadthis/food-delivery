-- +goose Up
CREATE TABLE IF NOT EXISTS products(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price FLOAT,
    description TEXT
);


-- +goose Down
DROP TABLE products;