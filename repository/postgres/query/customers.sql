-- name: InsertCustomer :one
INSERT INTO customers (first_name, last_name, tenant_id, last_modified_at, last_modified_by)
    VALUES ($1, $2, $3, $4, $5)
RETURNING
    id, first_name, last_name, tenant_id, last_modified_at, last_modified_by;

-- name: GetCustomers :many
SELECT
    id,
    first_name,
    last_name,
    tenant_id,
    last_modified_at,
    last_modified_by
FROM
    customers
WHERE
    deleted_at IS NULL
    AND tenant_id = $1;

-- name: GetCustomer :one
SELECT
    id,
    first_name,
    last_name,
    tenant_id,
    last_modified_at,
    last_modified_by
FROM
    customers
WHERE
    deleted_at IS NULL
    AND tenant_id = $1
    AND id = $2;


-- name: UpdateCustomer :one
UPDATE
    customers
SET
    first_name = $3,
    last_name = $4,
    last_modified_at = $5,
    last_modified_by = $6
WHERE
    deleted_at IS NULL
    AND tenant_id = $1
    AND id = $2
RETURNING
    id,
    first_name,
    last_name,
    tenant_id,
    last_modified_at,
    last_modified_by;

-- name: DeleteCustomer :exec
UPDATE
    customers
SET
    deleted_at = $3
WHERE
    deleted_at IS NULL
    AND tenant_id = $1
    AND id = $2;

