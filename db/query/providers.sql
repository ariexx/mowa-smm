-- name: GetProviders :many
SELECT id, name, link, api_key, api_id FROM providers WHERE status = '1';

-- name: GetProviderByID :one
SELECT id, name, link, api_key, api_id FROM providers WHERE id = ?;

-- name: CreateProvider :exec
INSERT INTO providers (name, link, api_key, api_id) VALUES (?, ?, ?, ?);