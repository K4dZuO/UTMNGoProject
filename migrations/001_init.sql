CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);


CREATE TABLE IF NOT EXISTS products (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    rate DOUBLE PRECISION DEFAULT 0,
    category_id INTEGER NOT NULL REFERENCES categories(id)
);


CREATE TABLE IF NOT EXISTS reviews (
    id UUID PRIMARY KEY,
    rate SMALLINT NOT NULL CHECK (rate >= 1 AND rate <= 5),
    product_id TEXT NOT NULL REFERENCES products(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);


-- Index to speed selecting reviews by product
CREATE INDEX IF NOT EXISTS idx_reviews_product_id ON reviews(product_id);
