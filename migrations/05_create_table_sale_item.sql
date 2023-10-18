-- Sale Item Table
CREATE TABLE "sale_item" (
    id SERIAL PRIMARY KEY NOT NULL,
    sale_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    subtotal FLOAT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(50) NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    updated_by VARCHAR(50) NOT NULL,
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(50)
);

ALTER TABLE sale_item ADD CONSTRAINT fk_sale FOREIGN KEY (sale_id) REFERENCES "sale" (id);
ALTER TABLE sale_item ADD CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES "product" (id);