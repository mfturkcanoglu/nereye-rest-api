-- +goose Up
-- +goose StatementBegin
ALTER TABLE restaurant 
ADD COLUMN is_available BOOLEAN,
ADD COLUMN available_at_start TIME,
ADD COLUMN available_at_end TIME,
ADD COLUMN weekend_available_at_start TIME,
ADD COLUMN weekend_available_at_end TIME;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE restaurant 
DROP COLUMN is_available,
DROP COLUMN available_at_start,
DROP COLUMN available_at_end,
DROP COLUMN weekend_available_at_start,
DROP COLUMN weekend_available_at_end;
-- +goose StatementEnd
