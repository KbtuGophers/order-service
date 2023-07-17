CREATE TABLE IF NOT EXISTS service_order.orders (
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    id          UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    customer_id VARCHAR NOT NULL,
    amount      NUMERIC NOT NULL,
    currency    VARCHAR NOT NULL,
    status      VARCHAR NOT NULL, -- pending/processing/(cancelled/completed)
    data        JSONB NOT NULL,
    billing_id  VARCHAR NULL
);

CREATE TABLE IF NOT EXISTS service_order.items (
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    id              UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    order_id        UUID NOT NULL,
    store_id        VARCHAR NOT NULL,
    product_id      VARCHAR NOT NULL,
    quantity        NUMERIC NOT NULL,
    price           NUMERIC NOT NULL,
    currency        VARCHAR NOT NULL,
    FOREIGN KEY (order_id) REFERENCES service_order.orders (id)
);

CREATE TABLE IF NOT EXISTS service_order.processes (
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    id              UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    account_id      VARCHAR NOT NULL,
    order_id        UUID NOT NULL,
    order_status    VARCHAR NOT NULL,
    stage           VARCHAR NOT NULL,
    task            VARCHAR NOT NULL,
    method          VARCHAR NOT NULL,
    state           VARCHAR NOT NULL, -- pending/processing/(completed/failed)
    correlation_id  VARCHAR NULL,
    data            JSONB NULL,
    FOREIGN KEY (order_id) REFERENCES service_order.orders (id)
);