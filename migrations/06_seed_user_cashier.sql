CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Insert the cashier user with a bcrypt-encrypted password and generated UUID user_id
INSERT INTO "user" (id, username, password, role, created_at, created_by, updated_at, updated_by)
VALUES (
    gen_random_uuid(),
    'cashier1',
    crypt('cashier1password', gen_salt('bf')),
    'cashier',
    current_timestamp,
    'cashier1',
    current_timestamp,
    'cashier1'
);
