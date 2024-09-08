-- name: InsertCustomerNote :one
INSERT INTO customer_notes(note, customer_id, last_modified_at, last_modified_by)
    VALUES ($1, $2, $3, $4)
RETURNING
    id, note, customer_id, last_modified_at, last_modified_by;

-- name: GetCustomerNotes :many
SELECT
    id,
    note,
    customer_id,
    last_modified_at,
    last_modified_by
FROM
    customer_notes
WHERE
    customer_id = $1;

-- name: GetCustomerNote :one
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
    AND id = $2;

-- name: UpdateCustomerNote :one
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
    last_modified_by;

-- name: DeleteCustomerNote :exec
DELETE FROM customer_notes
WHERE customer_id = $1
    AND id = $2;

