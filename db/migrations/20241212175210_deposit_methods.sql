-- +goose Up
-- +goose StatementBegin
CREATE TABLE deposit_methods (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    deposit_category_id BIGINT NOT NULL,
    name VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    min_deposit BIGINT UNSIGNED NOT NULL,
    max_deposit BIGINT UNSIGNED NOT NULL,
    rate double NOT NULL,
    status VARCHAR(15) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    version INT NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE deposit_methods;
-- +goose StatementEnd
