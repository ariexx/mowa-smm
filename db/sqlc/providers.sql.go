// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: providers.sql

package db

import (
	"context"
	"database/sql"
)

const createProvider = `-- name: CreateProvider :exec
INSERT INTO providers (name, link, api_key, api_id) VALUES (?, ?, ?, ?)
`

type CreateProviderParams struct {
	Name   string         `json:"name"`
	Link   string         `json:"link"`
	ApiKey string         `json:"api_key"`
	ApiID  sql.NullString `json:"api_id"`
}

func (q *Queries) CreateProvider(ctx context.Context, arg CreateProviderParams) error {
	_, err := q.exec(ctx, q.createProviderStmt, createProvider,
		arg.Name,
		arg.Link,
		arg.ApiKey,
		arg.ApiID,
	)
	return err
}

const getProviderByID = `-- name: GetProviderByID :one
SELECT id, name, link, api_key, api_id FROM providers WHERE id = ?
`

type GetProviderByIDRow struct {
	ID     uint64         `json:"id"`
	Name   string         `json:"name"`
	Link   string         `json:"link"`
	ApiKey string         `json:"api_key"`
	ApiID  sql.NullString `json:"api_id"`
}

func (q *Queries) GetProviderByID(ctx context.Context, id uint64) (GetProviderByIDRow, error) {
	row := q.queryRow(ctx, q.getProviderByIDStmt, getProviderByID, id)
	var i GetProviderByIDRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Link,
		&i.ApiKey,
		&i.ApiID,
	)
	return i, err
}

const getProviders = `-- name: GetProviders :many
SELECT id, name, link, api_key, api_id FROM providers WHERE status = '1'
`

type GetProvidersRow struct {
	ID     uint64         `json:"id"`
	Name   string         `json:"name"`
	Link   string         `json:"link"`
	ApiKey string         `json:"api_key"`
	ApiID  sql.NullString `json:"api_id"`
}

func (q *Queries) GetProviders(ctx context.Context) ([]GetProvidersRow, error) {
	rows, err := q.query(ctx, q.getProvidersStmt, getProviders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProvidersRow
	for rows.Next() {
		var i GetProvidersRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Link,
			&i.ApiKey,
			&i.ApiID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}