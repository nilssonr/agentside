-- name: InsertCustomerPhoneNumber :one
INSERT INTO customer_phone_numbers (phone_number, type, customer_id, last_modified_at, last_modified_by)
    VALUES ($1, $2, $3, $4, $5)
RETURNING
    id, phone_number, type, customer_id, last_modified_at, last_modified_by;

-- name: GetCustomerPhoneNumbers :many
SELECT
    id,
    phone_number,
    type,
    customer_id,
    last_modified_at,
    last_modified_by
FROM
    customer_phone_numbers
WHERE
    customer_id = $1;

-- name: GetCustomerPhoneNumber :one
SELECT
    id,
    phone_number,
    type,
    customer_id,
    last_modified_at,
    last_modified_by
FROM
    customer_phone_numbers
WHERE
    customer_id = $1
    AND id = $2;

-- name: UpdateCustomerPhoneNumber :one
UPDATE
    customer_phone_numbers
SET
    phone_number = $3,
    type = $4,
    last_modified_at = $5,
    last_modified_by = $6
WHERE
    customer_id = $1
    AND id = $2
RETURNING
    id,
    phone_number,
    type,
    customer_id,
    last_modified_at,
    last_modified_by;

-- name: DeleteCustomerPhoneNumber :exec
DELETE FROM customer_phone_numbers
WHERE customer_id = $1
    AND id = $2;

