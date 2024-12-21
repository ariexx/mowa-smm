-- name: InsertBatchProviders :exec
INSERT INTO services (name, min, max, price, margin, service_category_id, provider_id, status, is_refill, is_cancelable) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetServicesByProviderID :many
SELECT id, name, min, max, price, margin, service_category_id, provider_id, status, is_refill, is_cancelable FROM services WHERE provider_id = ?;

-- name: GetServiceByID :one
SELECT id, name, min, max, price, margin, service_category_id, provider_id, status, is_refill, is_cancelable FROM services WHERE id = ?;

-- name: UpdateServiceByID :exec
UPDATE services SET name = ?, min = ?, max = ?, price = ?, margin = ?, service_category_id = ?, provider_id = ?, status = ?, is_refill = ?, is_cancelable = ? WHERE id = ?;

-- name: DeleteServiceByID :exec
UPDATE services SET deleted_at = NOW() WHERE id = ?;

-- name: GetServices :many
SELECT id, name, min, max, price, margin, service_category_id, provider_id, status, is_refill, is_cancelable FROM services WHERE deleted_at IS NULL;

