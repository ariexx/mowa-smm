-- +goose Up

-- +goose StatementBegin
CREATE TABLE wallet_act_types (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO wallet_act_types (name, description) VALUES
    ('credit', 'Credit to wallet account'),
    ('debit', 'Debit from wallet account');
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE wallet_journals (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    act_type_id BIGINT UNSIGNED NOT NULL,
    user_id integer UNSIGNED NOT NULL,
    amount INT NOT NULL,
    balance INT NOT NULL,
    trx_type VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (act_type_id) REFERENCES wallet_act_types(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE wallet_journals;
-- +goose StatementEnd

-- +goose StatementBegin
DELETE FROM wallet_act_types WHERE name IN ('credit', 'debit');
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE wallet_act_types;
-- +goose StatementEnd
