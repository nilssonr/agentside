-- name: InsertCustomerAddress :one
INSERT INTO customer_addresses(street_address, type, state, zip_code, country, customer_id, last_modified_at, last_modified_by)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING
    id, type, street_address, state, zip_code, country, customer_id, last_modified_at, last_modified_by;

-- name: GetCustomerAddresses :many
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
    customer_id = $1;

-- name: GetCustomerAddress :one
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
    AND id = $2;

-- name: UpdateCustomerAddress :one
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
    last_modified_by;

-- name: DeleteCustomerAddress :exec
DELETE FROM customer_addresses
WHERE customer_id = $1
    AND id = $2;

