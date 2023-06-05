-- order 
CREATE TABLE IF NOT EXISTS 'order' (
    id UUID NOT NULL PRIMARY KEY,
    car_id UUID NOT NULL,--
    client_id UUID NOT NULL,--
    tarif_id UUID NOT NULL,--
    total_price DOUBLE PRECISION NOT NULL,
    paid_price DOUBLE PRECISION NOT NULL,
    day_count INTEGER,
    start_date DATE, 
    discount DOUBLE PRECISION,
    order_number VARCHAR,
    status BOOLEAN,
    mileage INTEGER,
    is_paid_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (client_id) REFERENCES client (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (tarif_id) REFERENCES tarif (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (car_id) REFERENCES tarif (id) ON DELETE CASCADE ON UPDATE CASCADE,
);

CREATE TABLE IF NOT EXISTS "order_history"(
    id UUID PRIMARY KEY,
    order_id UUID NOT NULL,
    client_id UUID NOT NULL,
    tarif_id UUID NOT NULL,
    car_id UUID NOT NULL,
    total_price DOUBLE PRECISION NOT NULL,
    comment VARCHAR(30),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

--client
CREATE TABLE IF NOT EXISTS "client" (
    id UUID PRIMARY KEY,
    first_name VARCHAR (50) NOT NULL,
    last_name VARCHAR (50) NOT NULL,
    middle_name VARCHAR(50),
    phone_number VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    photo VARCHAR(50),
    is_blocked BOOLEAN NOT NULL
);

--payment
CREATE TABLE IF NOT EXISTS "payment"(
    id UUID PRIMARY KEY,
    order_id UUID NOT NULL,
    payment_type VARCHAR (50),
    total_price DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

--migrations
CREATE TABLE IF NOT EXISTS "migrations"(
    dirty boolean NOT NULL,
    version bigint NOT NULL,
);

--mechanic
CREATE TABLE IF NOT EXISTS "mechanic"(
    id UUID PRIMARY KEY,
    fullname VARCHAR(100),
    phone_number VARCHAR(100),
    photo VARCHAR(100)
)

--car 
CREATE TABLE IF NOT EXISTS "models" (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
);

CREATE TABLE IF NOT EXISTS "car" (
    id UUID PRIMARY KEY,
    state_number VARCHAR(2) NOT NULL,
    tarif_id UUID NOT NULL,
    model_id UUID NOT NULL,
    status VARCHAR NOT NULL
    FOREIGN KEY (tarif_id) REFERENCES tarif (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (model_id) REFERENCES model (id) ON DELETE CASCADE ON UPDATE CASCADE,
);

CREATE TABLE IF NOT EXISTS "tarif"(
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    price DOUBLE PRECISION,
);

CREATE TABLE IF NOT EXISTS "give_car"(
    order_id UUID NOT NULL,
    miliage INTEGER,
    fuel VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    comment VARCHAR(100)
);

-- CREATE TABLE IF NOT EXISTS "car_activity"(
--     id UUID PRIMARY KEY,
--     car_id UUID  NOT NULL,
--     data DATE NOT NULL,
--     status VARCHAR(30)
-- );

CREATE TABLE IF NOT EXISTS "reference_car"(
    order_id UUID NOT NULL,
    miliage INTEGER,
    fuel VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    comment VARCHAR(100)
);