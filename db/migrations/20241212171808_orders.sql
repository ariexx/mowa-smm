-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    id_provider VARCHAR(50) NOT NULL,
    provider_id BIGINT UNSIGNED NOT NULL,
    user_id integer UNSIGNED NOT NULL,
    target VARCHAR(250) NOT NULL,
    total BIGINT UNSIGNED NOT NULL,
    beginning_value BIGINT UNSIGNED NULL,
    partial BIGINT UNSIGNED NULL,
    status VARCHAR(15) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    version INT NOT NULL DEFAULT 1,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX idx_user_id ON orders(user_id);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX idx_id_provider ON orders(id_provider);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_user_id ON orders;
-- +goose StatementEnd

-- +goose StatementBegin
DROP INDEX idx_id_provider ON orders;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
