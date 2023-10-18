-- Customer Table
CREATE TABLE "customer" (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    contact VARCHAR(50) NOT NULL,
    address TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(50) NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    updated_by VARCHAR(50) NOT NULL,
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(50)
);