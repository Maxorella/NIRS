CREATE TABLE "order" (
    order_id SERIAL PRIMARY KEY,
    user_id INT,
    order_date TIMESTAMP,
    address VARCHAR(128),
    total INT DEFAULT 0,
    FOREIGN KEY (user_id) REFERENCES "user"(user_id)
);