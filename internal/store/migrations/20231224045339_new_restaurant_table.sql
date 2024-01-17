-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS restaurant (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    photo_url TEXT,
    sign_name TEXT,
    address_id uuid,
    customer_id uuid,
    created_at TIMESTAMP DEFAULT (now() AT TIME ZONE 'utc') NOT NULL,
    updated_at TIMESTAMP DEFAULT (now() AT TIME ZONE 'utc') NOT NULL,
    deleted_at TIMESTAMP,
    deleted BOOLEAN,
    CONSTRAINT fk_customer FOREIGN KEY(customer_id) REFERENCES customer(id),
    CONSTRAINT fk_address FOREIGN KEY(address_id) REFERENCES address(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE restaurant;
-- +goose StatementEnd