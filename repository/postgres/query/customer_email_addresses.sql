-- name: InsertCustomerEmailAddress :one
INSERT INTO customer_email_addresses(email_address, type, customer_id, last_modified_at, last_modified_by)
    VALUES ($1, $2, $3, $4, $5)
RETURNING
    id, email_address, type, customer_id, last_modified_at, last_modified_by;

-- name: GetCustomerEmailAddresses :many
SELECT
    id,
    email_address,
    type,
    customer_id,
    last_modified_at,
    last_modified_by
FROM
    customer_email_addresses
WHERE
    customer_id = $1;

-- name: GetCustomerEmailAddress :one
SELECT
    id,
    email_address,
    type,
    customer_id,
    last_modified_at,
    last_modified_by
FROM
    customer_email_addresses
WHERE
    customer_id = $1
    AND id = $2;

-- name: UpdateCustomerEmailAddress :one
UPDATE
    customer_email_addresses
SET
    email_address = $3,
    type = $4,
    last_modified_at = $5,
    last_modified_by = $6
WHERE
    customer_id = $1
    AND id = $2
RETURNING
    id,
    email_address,
    type,
    customer_id,
    last_modified_at,
    last_modified_by;

-- name: DeleteCustomerEmailAddress :exec
DELETE FROM customer_email_addresses
WHERE customer_id = $1
    AND id = $2;

