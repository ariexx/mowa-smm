-- +goose Up
-- +goose StatementBegin
CREATE TABLE ticket_categories (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(15) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    version INT NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO ticket_categories (name, description, status) VALUES ('General', 'General ticket category', 'active');
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO ticket_categories (name, description, status) VALUES ('Technical', 'Technical ticket category', 'active');
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE tickets (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id integer UNSIGNED NOT NULL,
    ticket_category_id BIGINT UNSIGNED NOT NULL,
    title VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(15) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    version INT NOT NULL DEFAULT 1,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (ticket_category_id) REFERENCES ticket_categories(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO tickets (user_id, ticket_category_id, title, description, status) VALUES (1, 1, 'General ticket', 'General ticket description', 'active');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tickets DROP FOREIGN KEY tickets_ibfk_1;
ALTER TABLE tickets DROP FOREIGN KEY tickets_ibfk_2;
DROP TABLE tickets;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE ticket_categories;
-- +goose StatementEnd
