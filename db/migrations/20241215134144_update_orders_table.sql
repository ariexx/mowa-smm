-- +goose Up
-- +goose StatementBegin
ALTER TABLE orders ADD COLUMN price DOUBLE DEFAULT 0 AFTER total;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE orders DROP COLUMN price;
-- +goose StatementEnd
