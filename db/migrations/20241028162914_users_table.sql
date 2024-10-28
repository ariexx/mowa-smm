-- +goose Up
-- +goose StatementBegin
    CREATE TABLE users (
        id integer unsigned NOT NULL AUTO_INCREMENT,
        user_status_id integer unsigned NOT NULL, # Foreign key to user_statuses table
        user_type_id integer unsigned NOT NULL, # Foreign key to user_types table
        full_name varchar(100) NOT NULL,
        email varchar(100) NOT NULL UNIQUE,
        password text NOT NULL,
        phone_number varchar(20) NOT NULL,
        role enum('admin', 'member') NOT NULL DEFAULT 'member',
        status enum('active', 'inactive') NOT NULL DEFAULT 'active',
        email_verified_at timestamp NULL,
        phone_number_verified_at timestamp NULL,
        api_key varchar(100) NOT NULL UNIQUE,
        created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        deleted_at timestamp NULL,
        version integer NOT NULL DEFAULT 1,
        PRIMARY KEY (id),
        FOREIGN KEY (user_status_id) REFERENCES user_statuses(id),
        FOREIGN KEY (user_type_id) REFERENCES user_types(id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
# Insert some default data
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO users (user_status_id, user_type_id, full_name, email, password, phone_number, role, status, api_key)
VALUES (1, 1, 'Admin User', 'mowa@admin.com', sha2('password', 256), '08012345678', 'admin', 'active', rand(1000000000));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DROP TABLE users;
-- +goose StatementEnd
