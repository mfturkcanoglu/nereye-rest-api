-- +goose Up
-- +goose StatementBegin
ALTER TABLE restaurant
ALTER COLUMN available_at_start TYPE TEXT
USING TO_CHAR(available_at_start, 'HH12:MM');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
