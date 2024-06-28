CREATE TABLE "product"(
    product_id SERIAL PRIMARY KEY,
    product_name  VARCHAR(128) UNIQUE,
    price INT,
    stock INT
);

CREATE INDEX idx_product_name ON "product"(product_name);