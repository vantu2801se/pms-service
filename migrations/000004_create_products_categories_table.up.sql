CREATE TABLE IF NOT EXISTS pms_products_categories (
    product_id INT NOT NULL,
    category_id INT NOT NULL,
    PRIMARY KEY (product_id, category_id),
    FOREIGN KEY (product_id) REFERENCES pms_products(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES pms_categories(id) ON DELETE CASCADE
);