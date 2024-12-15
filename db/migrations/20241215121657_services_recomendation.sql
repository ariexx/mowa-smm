-- +goose Up
-- +goose StatementBegin
CREATE TABLE service_recomendations (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    service_id BIGINT UNSIGNED NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    version INT NOT NULL DEFAULT 1,
    PRIMARY KEY (id),
    FOREIGN KEY (service_id) REFERENCES services(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE service_recomendations DROP FOREIGN KEY service_recomendations_ibfk_1;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE service_recomendations;
-- +goose StatementEnd
