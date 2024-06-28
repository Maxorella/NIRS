CREATE TABLE order_detail (
    order_detail_id SERIAL PRIMARY KEY,
    order_id INT,
    product_id INT,
    quantity INT,
    price INT,
    FOREIGN KEY (order_id) REFERENCES order(order_id),
    FOREIGN KEY (product_id) REFERENCES product(product_id)
);

CREATE OR REPLACE FUNCTION update_order_total()
RETURNS TRIGGER AS $$
BEGIN
UPDATE order
SET total = total + NEW.quantity * NEW.price
WHERE order_id = NEW.order_id;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_insert_order_detail
    AFTER INSERT ON order_detail
    FOR EACH ROW
    EXECUTE FUNCTION update_order_total();