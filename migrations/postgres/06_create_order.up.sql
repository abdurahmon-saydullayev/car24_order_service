CREATE TABLE IF NOT EXISTS "order" (
    id UUID NOT NULL PRIMARY KEY,
    car_id UUID NOT NULL,
    client_id UUID NOT NULL,
    tarif_id UUID NOT NULL,
    total_price DOUBLE PRECISION NOT NULL,
    paid_price DOUBLE PRECISION DEFAULT 0,
    day_count INTEGER,
    start_date DATE, 
    discount DOUBLE PRECISION,
    order_number VARCHAR,
    status BOOLEAN,
    mileage INTEGER,
    is_paid_date TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (tarif_id) REFERENCES tarif (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (car_id) REFERENCES car (id) ON DELETE CASCADE ON UPDATE CASCADE
);
