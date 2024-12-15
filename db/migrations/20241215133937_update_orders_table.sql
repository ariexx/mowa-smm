-- +goose Up
-- +goose StatementBegin
ALTER TABLE orders ADD COLUMN name VARCHAR(255) DEFAULT NULL AFTER id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE orders DROP COLUMN name;
-- +goose StatementEnd
