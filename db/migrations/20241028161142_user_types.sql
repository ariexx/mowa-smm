-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_types (
                            id integer unsigned NOT NULL AUTO_INCREMENT,
                            name varchar(50) NOT NULL UNIQUE,
                            description text NOT NULL,
                            created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            deleted_at timestamp NULL,
                            version integer NOT NULL DEFAULT 1,
                            PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO user_types (name, description) VALUES ('admin', 'Admin user type');
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO user_types (name, description) VALUES ('member', 'Member user type');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_types;
-- +goose StatementEnd