CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

ALTER TABLE products 
DROP COLUMN category;

ALTER TABLE products
ADD COLUMN category_id INT;

ALTER TABLE products
ADD CONSTRAINT fk_products_categories
FOREIGN KEY (category_id) REFERENCES categories (id);

