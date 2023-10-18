-- Sale Table
CREATE TABLE "sale" (
    id SERIAL PRIMARY KEY NOT NULL,
    user_id VARCHAR(36) UNIQUE NOT NULL,
    customer_id VARCHAR(36) UNIQUE NOT NULL,
    sale_date TIMESTAMP,
    total_amount INT,
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(50) NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    updated_by VARCHAR(50) NOT NULL,
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(50)
);

-- Add foreign key constraints
ALTER TABLE "sale" ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES "user" (id);
ALTER TABLE "sale" ADD CONSTRAINT fk_customer FOREIGN KEY (customer_id) REFERENCES "customer" (id);