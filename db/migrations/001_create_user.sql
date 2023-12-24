-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    phone_number TEXT UNIQUE NOT NULL,
    full_name TEXT,
    surname TEXT,
    verified BOOLEAN,
    enabled BOOLEAN,
    active BOOLEAN,
    created_at TIMESTAMP DEFAULT (now() AT TIME ZONE 'utc') NOT NULL,
    updated_at TIMESTAMP DEFAULT (now() AT TIME ZONE 'utc') NOT NULL,
    deleted_at TIMESTAMP,
    deleted BOOLEAN
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd