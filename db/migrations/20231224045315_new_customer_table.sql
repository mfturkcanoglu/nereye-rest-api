-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS customer(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    company_name TEXT,
    customer_type TEXT,
    address_id uuid,
    user_id uuid NOT NULL,
    created_at TIMESTAMP DEFAULT (now() AT TIME ZONE 'utc') NOT NULL,
    updated_at TIMESTAMP DEFAULT (now() AT TIME ZONE 'utc') NOT NULL,
    deleted_at TIMESTAMP,
    deleted BOOLEAN,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_address FOREIGN KEY(address_id) REFERENCES address(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE customer;
-- +goose StatementEnd