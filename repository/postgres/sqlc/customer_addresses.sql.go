// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: customer_addresses.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteCustomerAddress = `-- name: DeleteCustomerAddress :exec
DELETE FROM customer_addresses
WHERE customer_id = $1
    AND id = $2
`

type DeleteCustomerAddressParams struct {
	CustomerID string
	ID         string
}

func (q *Queries) DeleteCustomerAddress(ctx context.Context, arg DeleteCustomerAddressParams) error {
	_, err := q.db.Exec(ctx, deleteCustomerAddress, arg.CustomerID, arg.ID)
	return err
}

const getCustomerAddress = `-- name: GetCustomerAddress :one
SELECT
    id,
    type,
    street_address,
    state,
    zip_code,
    country,
    customer_id,
    last_modified_at,
    last_modified_by
FROM
    customer_addresses
WHERE
    customer_id = $1
    AND id = $2
`

type GetCustomerAddressParams struct {
	CustomerID string
	ID         string
}

func (q *Queries) GetCustomerAddress(ctx context.Context, arg GetCustomerAddressParams) (CustomerAddress, error) {
	row := q.db.QueryRow(ctx, getCustomerAddress, arg.CustomerID, arg.ID)
	var i CustomerAddress
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.StreetAddress,
		&i.State,
		&i.ZipCode,
		&i.Country,
		&i.CustomerID,
		&i.LastModifiedAt,
		&i.LastModifiedBy,
	)
	return i, err
}

const getCustomerAddresses = `-- name: GetCustomerAddresses :many
SELECT
    id,
    type,
    street_address,
    state,
    zip_code,
    country,
    customer_id,
    last_modified_at,
    last_modified_by
FROM
    customer_addresses
WHERE
    customer_id = $1
`

func (q *Queries) GetCustomerAddresses(ctx context.Context, customerID string) ([]CustomerAddress, error) {
	rows, err := q.db.Query(ctx, getCustomerAddresses, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CustomerAddress
	for rows.Next() {
		var i CustomerAddress
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.StreetAddress,
			&i.State,
			&i.ZipCode,
			&i.Country,
			&i.CustomerID,
			&i.LastModifiedAt,
			&i.LastModifiedBy,
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

const insertCustomerAddress = `-- name: InsertCustomerAddress :one
INSERT INTO customer_addresses(street_address, type, state, zip_code, country, customer_id, last_modified_at, last_modified_by)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING
    id, type, street_address, state, zip_code, country, customer_id, last_modified_at, last_modified_by
`

type InsertCustomerAddressParams struct {
	StreetAddress  string
	Type           string
	State          pgtype.Text
	ZipCode        pgtype.Text
	Country        pgtype.Text
	CustomerID     string
	LastModifiedAt pgtype.Timestamptz
	LastModifiedBy string
}

func (q *Queries) InsertCustomerAddress(ctx context.Context, arg InsertCustomerAddressParams) (CustomerAddress, error) {
	row := q.db.QueryRow(ctx, insertCustomerAddress,
		arg.StreetAddress,
		arg.Type,
		arg.State,
		arg.ZipCode,
		arg.Country,
		arg.CustomerID,
		arg.LastModifiedAt,
		arg.LastModifiedBy,
	)
	var i CustomerAddress
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.StreetAddress,
		&i.State,
		&i.ZipCode,
		&i.Country,
		&i.CustomerID,
		&i.LastModifiedAt,
		&i.LastModifiedBy,
	)
	return i, err
}

const updateCustomerAddress = `-- name: UpdateCustomerAddress :one
UPDATE
    customer_addresses
SET
    type = $3,
    street_address = $4,
    state = $5,
    zip_code = $6,
    country = $7,
    last_modified_at = $8,
    last_modified_by = $9
WHERE
    customer_id = $1
    AND id = $2
RETURNING
    id,
    type,
    street_address,
    state,
    zip_code,
    country,
    customer_id,
    last_modified_at,
    last_modified_by
`

type UpdateCustomerAddressParams struct {
	CustomerID     string
	ID             string
	Type           string
	StreetAddress  string
	State          pgtype.Text
	ZipCode        pgtype.Text
	Country        pgtype.Text
	LastModifiedAt pgtype.Timestamptz
	LastModifiedBy string
}

func (q *Queries) UpdateCustomerAddress(ctx context.Context, arg UpdateCustomerAddressParams) (CustomerAddress, error) {
	row := q.db.QueryRow(ctx, updateCustomerAddress,
		arg.CustomerID,
		arg.ID,
		arg.Type,
		arg.StreetAddress,
		arg.State,
		arg.ZipCode,
		arg.Country,
		arg.LastModifiedAt,
		arg.LastModifiedBy,
	)
	var i CustomerAddress
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.StreetAddress,
		&i.State,
		&i.ZipCode,
		&i.Country,
		&i.CustomerID,
		&i.LastModifiedAt,
		&i.LastModifiedBy,
	)
	return i, err
}
