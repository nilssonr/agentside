-- name: InsertCustomerAddress :one
INSERT INTO customer_addresses(street_address, state, zip_code, country, customer_id, last_modified_at, last_modified_by)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING
    id, street_address, state, zip_code, country, customer_id, last_modified_at, last_modified_by;

-- name: GetCustomerAddresses :many
SELECT
    id,
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
    street_address = $3,
    state = $4,
    zip_code = $5,
    country = $6,
    last_modified_at = $7,
    last_modified_by = $8
WHERE
    customer_id = $1
    AND id = $2
RETURNING
    id,
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

