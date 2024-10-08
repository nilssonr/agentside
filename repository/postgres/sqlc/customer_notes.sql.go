// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: customer_notes.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteCustomerNote = `-- name: DeleteCustomerNote :exec
DELETE FROM customer_notes
WHERE customer_id = $1
    AND id = $2
`

type DeleteCustomerNoteParams struct {
	CustomerID string
	ID         string
}

func (q *Queries) DeleteCustomerNote(ctx context.Context, arg DeleteCustomerNoteParams) error {
	_, err := q.db.Exec(ctx, deleteCustomerNote, arg.CustomerID, arg.ID)
	return err
}

const getCustomerNote = `-- name: GetCustomerNote :one
SELECT
    id,
    note,
    customer_id,
    last_modified_at,
    last_modified_by
FROM
    customer_notes
WHERE
    customer_id = $1
    AND id = $2
`

type GetCustomerNoteParams struct {
	CustomerID string
	ID         string
}

func (q *Queries) GetCustomerNote(ctx context.Context, arg GetCustomerNoteParams) (CustomerNote, error) {
	row := q.db.QueryRow(ctx, getCustomerNote, arg.CustomerID, arg.ID)
	var i CustomerNote
	err := row.Scan(
		&i.ID,
		&i.Note,
		&i.CustomerID,
		&i.LastModifiedAt,
		&i.LastModifiedBy,
	)
	return i, err
}

const getCustomerNotes = `-- name: GetCustomerNotes :many
SELECT
    id,
    note,
    customer_id,
    last_modified_at,
    last_modified_by
FROM
    customer_notes
WHERE
    customer_id = $1
`

func (q *Queries) GetCustomerNotes(ctx context.Context, customerID string) ([]CustomerNote, error) {
	rows, err := q.db.Query(ctx, getCustomerNotes, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CustomerNote
	for rows.Next() {
		var i CustomerNote
		if err := rows.Scan(
			&i.ID,
			&i.Note,
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

const insertCustomerNote = `-- name: InsertCustomerNote :one
INSERT INTO customer_notes(note, customer_id, last_modified_at, last_modified_by)
    VALUES ($1, $2, $3, $4)
RETURNING
    id, note, customer_id, last_modified_at, last_modified_by
`

type InsertCustomerNoteParams struct {
	Note           string
	CustomerID     string
	LastModifiedAt pgtype.Timestamptz
	LastModifiedBy string
}

func (q *Queries) InsertCustomerNote(ctx context.Context, arg InsertCustomerNoteParams) (CustomerNote, error) {
	row := q.db.QueryRow(ctx, insertCustomerNote,
		arg.Note,
		arg.CustomerID,
		arg.LastModifiedAt,
		arg.LastModifiedBy,
	)
	var i CustomerNote
	err := row.Scan(
		&i.ID,
		&i.Note,
		&i.CustomerID,
		&i.LastModifiedAt,
		&i.LastModifiedBy,
	)
	return i, err
}

const updateCustomerNote = `-- name: UpdateCustomerNote :one
UPDATE
    customer_notes
SET
    note = $3,
    last_modified_at = $4,
    last_modified_by = $5
WHERE
    customer_id = $1
    AND id = $2
RETURNING
    id,
    note,
    customer_id,
    last_modified_at,
    last_modified_by
`

type UpdateCustomerNoteParams struct {
	CustomerID     string
	ID             string
	Note           string
	LastModifiedAt pgtype.Timestamptz
	LastModifiedBy string
}

func (q *Queries) UpdateCustomerNote(ctx context.Context, arg UpdateCustomerNoteParams) (CustomerNote, error) {
	row := q.db.QueryRow(ctx, updateCustomerNote,
		arg.CustomerID,
		arg.ID,
		arg.Note,
		arg.LastModifiedAt,
		arg.LastModifiedBy,
	)
	var i CustomerNote
	err := row.Scan(
		&i.ID,
		&i.Note,
		&i.CustomerID,
		&i.LastModifiedAt,
		&i.LastModifiedBy,
	)
	return i, err
}
