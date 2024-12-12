-- +goose Up
-- +goose StatementBegin
    CREATE TABLE user_statuses (
        id integer unsigned NOT NULL AUTO_INCREMENT,
        name varchar(50) NOT NULL UNIQUE,
        description text NOT NULL,
        created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        deleted_at timestamp NULL,
        version integer NOT NULL DEFAULT 1,
        PRIMARY KEY (id),
        FOREIGN KEY (id) REFERENCES users(user_status_id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose StatementBegin
    INSERT INTO user_statuses (name, description) VALUES ('active', 'Active user status');
-- +goose StatementEnd

-- +goose StatementBegin
    INSERT INTO user_statuses (name, description) VALUES ('inactive', 'Inactive user status');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DROP TABLE user_statuses;
-- +goose StatementEnd
