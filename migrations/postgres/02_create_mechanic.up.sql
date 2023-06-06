CREATE TABLE IF NOT EXISTS "mechanic"(
    id UUID PRIMARY KEY,
    fullname VARCHAR(100),
    phone_number VARCHAR(100),
    photo VARCHAR(100),
    price_per_hour NUMERIC,
    status BOOLEAN DEFAULT FALSE
);