CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    user_id UUID NOT NULL,
    product_id UUID[] NOT NULL,
    pricing VARCHAR NOT NULL,
    status VARCHAR NOT NULL CHECK (status IN ('new', 'in_progress', 'done')),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


