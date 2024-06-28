CREATE TABLE order (
    order_id SERIAL PRIMARY KEY,
    user_id INT,
    order_date TIMESTAMP,
    address VARCHAR(100),
    total INT,
    FOREIGN KEY (user_id) REFERENCES user(user_id)
);