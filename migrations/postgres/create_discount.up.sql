CREATE TYPE discount_type AS ENUM ('fix', 'percentage');

CREATE TABLE IF NOT EXISTS "discount"(
    id UUID PRIMARY KEY,
    discount_type discount_type,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);