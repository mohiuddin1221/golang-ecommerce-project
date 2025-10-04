CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,
    title    VARCHAR(100) NOT NULL,
    description     TEXT,
    price DOUBLE PRECISION NOT NULL,
    created_at    TIMESTAMP DEFAULT NOW(),
    updated_at    TIMESTAMP DEFAULT NOW()
);
