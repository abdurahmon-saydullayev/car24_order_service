CREATE TABLE IF NOT EXISTS "tarif"(
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    model_id UUID NOT NULL,
    price_per_hour DOUBLE PRECISION
);