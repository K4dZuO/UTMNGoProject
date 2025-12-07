-- +goose Down
DROP INDEX IF EXISTS idx_reviews_product_id;
DROP TABLE IF EXISTS reviews;
DROP TABLE IF EXISTS products;
