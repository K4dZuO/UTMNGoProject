-- Categories
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- Products
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    rate DOUBLE PRECISION NOT NULL DEFAULT 0,
    category_id INT NOT NULL,
    CONSTRAINT fk_products_categories
        FOREIGN KEY (category_id)
        REFERENCES categories(id)
        ON DELETE RESTRICT
);

-- Reviews
CREATE TABLE reviews (
    id UUID PRIMARY KEY,
    product_id INT NOT NULL,
    rate SMALLINT NOT NULL CHECK (rate >= 1 AND rate <= 5),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_reviews_products
        FOREIGN KEY (product_id)
        REFERENCES products(id)
        ON DELETE CASCADE
);

-- Indexes
CREATE INDEX idx_reviews_product_id
    ON reviews(product_id);

CREATE INDEX idx_products_category_id
    ON products(category_id);

CREATE INDEX idx_products_rate_desc
    ON products(rate DESC);
