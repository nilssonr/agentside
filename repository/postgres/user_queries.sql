-- name: CreateUser :one
INSERT INTO users (first_name, last_name, email_address, tenant_id, last_modified_at, last_modified_by)
    VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
    id, first_name, last_name, email_address, tenant_id, last_modified_at, last_modified_by;

-- name: GetUsers :many
SELECT
    id,
    first_name,
    last_name,
    email_address,
    tenant_id,
    last_modified_at,
    last_modified_by
FROM
    users
WHERE
    deleted_at IS NULL
    AND tenant_id = $1;

-- name: GetUser :one
SELECT
    id,
    first_name,
    last_name,
    email_address,
    tenant_id,
    last_modified_at,
    last_modified_by
FROM
    users
WHERE
    deleted_at IS NULL
    AND tenant_id = $1
    AND id = $2;

-- name: GetUserByEmailAddress :one
SELECT
    id,
    first_name,
    last_name,
    email_address,
    tenant_id,
    last_modified_at,
    last_modified_by
FROM
    users
WHERE
    deleted_at IS NULL
    AND email_address = $1;


-- name: UpdateUser :one
UPDATE
    users
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
    email_address,
    tenant_id,
    last_modified_at,
    last_modified_by;

-- name: DeleteUser :exec
UPDATE
    users
SET
    deleted_at = $3
WHERE
    deleted_at IS NULL
    AND tenant_id = $1
    AND id = $2;
