-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users WHERE id = ? LIMIT 1;

-- name: CreateUser :exec
INSERT INTO users (id, user_status_id, user_type_id, full_name, email, password, phone_number, role, status, email_verified_at, phone_number_verified_at, api_key, created_at, updated_at, deleted_at, version)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateUser :exec
UPDATE users SET user_status_id = ?, user_type_id = ?, full_name = ?, email = ?, password = ?, phone_number = ?, role = ?, status = ?, email_verified_at = ?, phone_number_verified_at = ?, api_key = ?, created_at = ?, updated_at = ?, deleted_at = ?, version = ? WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ? LIMIT 1;

-- name: GetUserByPhoneNumber :one
SELECT * FROM users WHERE phone_number = ? LIMIT 1;

-- name: GetUserByEmailAndPassword :one
SELECT * FROM users WHERE email = ? AND password = ? LIMIT 1;

-- name: IsUserEmailExists :one
SELECT EXISTS(SELECT 1 FROM users WHERE email = ?) AS email_exists;

-- name: IsUserPhoneNumberExists :one
SELECT EXISTS(SELECT 1 FROM users WHERE phone_number = ?) AS phone_number_exists;

-- name: IsUserDeleted :one
SELECT deleted_at IS NOT NULL AS is_deleted FROM users WHERE id = ?;

-- name: IsUserEmailVerified :one
SELECT email_verified_at IS NOT NULL AS is_email_verified FROM users WHERE id = ?;

-- name: IsUserPhoneNumberVerified :one
SELECT phone_number_verified_at IS NOT NULL AS is_phone_number_verified FROM users WHERE id = ?;

-- name: IsUserActive :one
SELECT status = 'active' AS is_active FROM users WHERE id = ?;

-- name: IsUserAdmin :one
SELECT role = 'admin' AS is_admin FROM users WHERE id = ?;

-- name: GetUserByIdAndVersion :one
SELECT * FROM users WHERE id = ? AND version = ? LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = ? LIMIT 1;
