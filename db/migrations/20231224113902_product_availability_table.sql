-- +goose Up
-- +goose StatementBegin
ALTER TABLE product ADD COLUMN is_available BOOLEAN;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE product DROP COLUMN is_available BOOLEAN;
-- +goose StatementEnd
