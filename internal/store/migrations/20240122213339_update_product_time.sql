-- +goose Up
-- +goose StatementBegin
ALTER TABLE product
ALTER COLUMN available_at_end TYPE TEXT
USING TO_CHAR(available_at_end, 'HH12:MM');

ALTER TABLE product
ALTER COLUMN available_at_start TYPE TEXT
USING TO_CHAR(available_at_start, 'HH12:MM');

ALTER TABLE product
DROP COLUMN is_available;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
