// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: auth_clients.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteAuthClient = `-- name: DeleteAuthClient :exec
UPDATE
    auth_clients
SET
    deleted_at = $3
WHERE
    deleted_at IS NULL
    AND tenant_id = $1
    AND id = $2
`

type DeleteAuthClientParams struct {
	TenantID  string
	ID        string
	DeletedAt pgtype.Timestamptz
}

func (q *Queries) DeleteAuthClient(ctx context.Context, arg DeleteAuthClientParams) error {
	_, err := q.db.Exec(ctx, deleteAuthClient, arg.TenantID, arg.ID, arg.DeletedAt)
	return err
}

const getAuthClient = `-- name: GetAuthClient :one
SELECT
    id,
    name,
    secret,
    last_modified_at,
    last_modified_by,
    tenant_id
FROM
    auth_clients
WHERE
    deleted_at IS NULL
    AND tenant_id = $1
    AND id = $2
`

type GetAuthClientParams struct {
	TenantID string
	ID       string
}

type GetAuthClientRow struct {
	ID             string
	Name           string
	Secret         string
	LastModifiedAt pgtype.Timestamptz
	LastModifiedBy string
	TenantID       string
}

func (q *Queries) GetAuthClient(ctx context.Context, arg GetAuthClientParams) (GetAuthClientRow, error) {
	row := q.db.QueryRow(ctx, getAuthClient, arg.TenantID, arg.ID)
	var i GetAuthClientRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Secret,
		&i.LastModifiedAt,
		&i.LastModifiedBy,
		&i.TenantID,
	)
	return i, err
}

const getAuthClients = `-- name: GetAuthClients :many
SELECT
    id,
    name,
    secret,
    last_modified_at,
    last_modified_by,
    tenant_id
FROM
    auth_clients
WHERE
    deleted_at IS NULL
    AND tenant_id = $1
`

type GetAuthClientsRow struct {
	ID             string
	Name           string
	Secret         string
	LastModifiedAt pgtype.Timestamptz
	LastModifiedBy string
	TenantID       string
}

func (q *Queries) GetAuthClients(ctx context.Context, tenantID string) ([]GetAuthClientsRow, error) {
	rows, err := q.db.Query(ctx, getAuthClients, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAuthClientsRow
	for rows.Next() {
		var i GetAuthClientsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Secret,
			&i.LastModifiedAt,
			&i.LastModifiedBy,
			&i.TenantID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertAuthClient = `-- name: InsertAuthClient :one
INSERT INTO auth_clients(name, secret, last_modified_at, last_modified_by, tenant_id)
    VALUES ($1, $2, $3, $4, $5)
RETURNING
    id, name, secret, last_modified_at, last_modified_by, tenant_id
`

type InsertAuthClientParams struct {
	Name           string
	Secret         string
	LastModifiedAt pgtype.Timestamptz
	LastModifiedBy string
	TenantID       string
}

type InsertAuthClientRow struct {
	ID             string
	Name           string
	Secret         string
	LastModifiedAt pgtype.Timestamptz
	LastModifiedBy string
	TenantID       string
}

func (q *Queries) InsertAuthClient(ctx context.Context, arg InsertAuthClientParams) (InsertAuthClientRow, error) {
	row := q.db.QueryRow(ctx, insertAuthClient,
		arg.Name,
		arg.Secret,
		arg.LastModifiedAt,
		arg.LastModifiedBy,
		arg.TenantID,
	)
	var i InsertAuthClientRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Secret,
		&i.LastModifiedAt,
		&i.LastModifiedBy,
		&i.TenantID,
	)
	return i, err
}

const updateAuthClient = `-- name: UpdateAuthClient :one
UPDATE
    auth_clients
SET
    name = $3,
    last_modified_at = $4,
    last_modified_by = $5
WHERE
    deleted_at IS NULL
    AND tenant_id = $1
    AND id = $2
RETURNING
    id,
    name,
    secret,
    last_modified_at,
    last_modified_by,
    tenant_id
`

type UpdateAuthClientParams struct {
	TenantID       string
	ID             string
	Name           string
	LastModifiedAt pgtype.Timestamptz
	LastModifiedBy string
}

type UpdateAuthClientRow struct {
	ID             string
	Name           string
	Secret         string
	LastModifiedAt pgtype.Timestamptz
	LastModifiedBy string
	TenantID       string
}

func (q *Queries) UpdateAuthClient(ctx context.Context, arg UpdateAuthClientParams) (UpdateAuthClientRow, error) {
	row := q.db.QueryRow(ctx, updateAuthClient,
		arg.TenantID,
		arg.ID,
		arg.Name,
		arg.LastModifiedAt,
		arg.LastModifiedBy,
	)
	var i UpdateAuthClientRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Secret,
		&i.LastModifiedAt,
		&i.LastModifiedBy,
		&i.TenantID,
	)
	return i, err
}
