-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS address (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    country TEXT,
    city TEXT,
    county TEXT,
    district TEXT,
    full_address TEXT,
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    created_at TIMESTAMP DEFAULT (now() AT TIME ZONE 'utc') NOT NULL,
    updated_at TIMESTAMP DEFAULT (now() AT TIME ZONE 'utc') NOT NULL,
    deleted_at TIMESTAMP,
    deleted BOOLEAN
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE address;
-- +goose StatementEnd