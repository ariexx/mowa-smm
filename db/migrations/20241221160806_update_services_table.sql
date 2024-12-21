-- +goose Up
-- +goose StatementBegin
ALTER TABLE services ADD COLUMN is_refill BOOLEAN DEFAULT FALSE;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE services ADD COLUMN is_cancelable BOOLEAN DEFAULT FALSE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE services DROP COLUMN is_refill;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE services DROP COLUMN is_cancelable;
-- +goose StatementEnd
