-- name: InsertTenant :one
INSERT INTO tenants(name, last_modified_at, last_modified_by)
    VALUES ($1, $2, $3)
RETURNING
    id, name, last_modified_at, last_modified_by;

-- name: GetTenants :many
SELECT
    id,
    name,
    last_modified_at,
    last_modified_by
FROM
    tenants
WHERE
    deleted_at IS NULL;

-- name: GetTenant :one
SELECT
    id,
    name,
    last_modified_at,
    last_modified_by
FROM
    tenants
WHERE
    deleted_at IS NULL
    AND id = $1;

-- name: UpdateTenant :one
UPDATE
    tenants
SET
    name = $2,
    last_modified_at = $3,
    last_modified_by = $4
WHERE
    deleted_at IS NULL
    AND id = $1
RETURNING
    id,
    name,
    last_modified_at,
    last_modified_by;

-- name: DeleteTenant :exec
UPDATE
    tenants
SET
    deleted_at = $2
WHERE
    deleted_at IS NULL
    AND id = $1;

