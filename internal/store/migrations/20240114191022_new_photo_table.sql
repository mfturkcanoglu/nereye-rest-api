-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS restaurant_photo (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT,
    url TEXT,
    parent_id uuid,
    photo_order SMALLINT,
    created_at TIMESTAMP DEFAULT (now() AT TIME ZONE 'utc') NOT NULL,
    updated_at TIMESTAMP DEFAULT (now() AT TIME ZONE 'utc') NOT NULL,
    deleted_at TIMESTAMP,
    deleted BOOLEAN,
    CONSTRAINT fk_restaurant FOREIGN KEY(parent_id) REFERENCES restaurant(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE restaurant_photo;
-- +goose StatementEnd