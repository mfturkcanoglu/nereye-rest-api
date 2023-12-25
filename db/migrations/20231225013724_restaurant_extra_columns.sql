-- +goose Up
-- +goose StatementBegin
ALTER TABLE restaurant 
ADD COLUMN about_us TEXT,
ADD COLUMN extra_info TEXT,
ADD COLUMN phone_number TEXT,
ADD COLUMN workplace_phone_number TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE restaurant 
DROP COLUMN about_us,
DROP COLUMN extra_info,
DROP COLUMN phone_number,
DROP COLUMN workplace_phone_number;
-- +goose StatementEnd
