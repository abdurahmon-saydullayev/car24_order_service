CREATE TYPE discount_type AS ENUM ('fix', 'percentage');

CREATE TABLE IF NOT EXISTS "discount"(
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    discount_type discount_type,
    discount_amount NUMERIC NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);