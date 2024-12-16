-- name: GetStatistics :one
SELECT COUNT(*) AS total_users FROM users;

-- name: GetAdmins :many
-- With Pagination
-- SELECT full_name, email, status, email_verified_at, created_at, updated_at FROM users WHERE role = 'admin';
SELECT full_name, email, status, email_verified_at, created_at, updated_at FROM users WHERE role = 'admin' LIMIT ? OFFSET ?;

-- name: GetLastOrders :many
-- Join Order with User
SELECT o.id, o.total, o.status, o.created_at, o.updated_at, u.full_name, u.email, o.name, o.price FROM orders o JOIN users u ON o.user_id = u.id ORDER BY o.created_at DESC LIMIT ?;
