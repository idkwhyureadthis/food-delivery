-- +goose Up
CREATE TABLE IF NOT EXISTS order_items(
    product_id BIGINT REFERENCES products(id),
    order_id INTEGER REFERENCES orders(id)
);


-- +goose Down
DROP TABLE order_items;