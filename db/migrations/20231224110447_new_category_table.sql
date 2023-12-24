-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS category (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    photo_url TEXT,
    category TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT (now() AT TIME ZONE 'utc') NOT NULL,
    updated_at TIMESTAMP DEFAULT (now() AT TIME ZONE 'utc') NOT NULL,
    deleted_at TIMESTAMP,
    deleted BOOLEAN
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE category;
-- +goose StatementEnd