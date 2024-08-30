-- +goose Up
CREATE TABLE IF NOT EXISTS orders(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(id),
    time TIMESTAMP,
    total_price FLOAT
);

-- +goose Down
DROP TABLE ORDERS;