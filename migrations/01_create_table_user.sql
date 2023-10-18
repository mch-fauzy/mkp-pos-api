-- Define the role enum type
CREATE TYPE role AS ENUM ('admin', 'cashier');

-- User Table
CREATE TABLE "user" (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    role role NOT NULL,
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(50) NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    updated_by VARCHAR(50) NOT NULL,
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(50)
);