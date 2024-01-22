-- +goose Up
-- +goose StatementBegin
ALTER TABLE restaurant
ALTER COLUMN available_at_end TYPE TEXT
USING TO_CHAR(available_at_end, 'HH12:MM');

ALTER COLUMN weekend_available_at_start TYPE TEXT
USING TO_CHAR(weekend_available_at_start, 'HH12:MM');

ALTER COLUMN weekend_available_at_end TYPE TEXT
USING TO_CHAR(weekend_available_at_end, 'HH12:MM');

DROP COLUMN is_available;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
