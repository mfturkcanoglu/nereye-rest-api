-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS product(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    product_name TEXT NOT NULL UNIQUE,
    photo_url TEXT,
    available_at_start TIME,
    available_at_end TIME,
    restaurant_id uuid NOT NULL,
    category_id uuid,
    created_at TIMESTAMP DEFAULT (now() AT TIME ZONE 'utc') NOT NULL,
    updated_at TIMESTAMP DEFAULT (now() AT TIME ZONE 'utc') NOT NULL,
    deleted_at TIMESTAMP,
    deleted BOOLEAN,
    CONSTRAINT fk_restaurant FOREIGN KEY(restaurant_id) REFERENCES restaurant(id),
    CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES category(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE product;
-- +goose StatementEnd