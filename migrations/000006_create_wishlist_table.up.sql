CREATE TABLE IF NOT EXISTS pms_wishlist (
    user_id INT NOT NULL,
    product_id INT NOT NULL,
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, product_id),
    FOREIGN KEY (user_id) REFERENCES pms_users(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES pms_products(id) ON DELETE CASCADE
);