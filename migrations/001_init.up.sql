-- +goose Up
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    rate DOUBLE PRECISION DEFAULT 0,
    category TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS reviews (
    id UUID PRIMARY KEY,
    rate SMALLINT NOT NULL CHECK (rate >= 1 AND rate <= 5),
    product_id INT NOT NULL REFERENCES products(id),
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_reviews_product_id
    ON reviews(product_id);
