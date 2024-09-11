-- name: InsertAuthClient :one
INSERT INTO auth_clients(name, secret, last_modified_at, last_modified_by, tenant_id)
    VALUES ($1, $2, $3, $4, $5)
RETURNING
    id, name, secret, last_modified_at, last_modified_by, tenant_id;

-- name: GetAuthClients :many
SELECT
    id,
    name,
    secret,
    last_modified_at,
    last_modified_by,
    tenant_id
FROM
    auth_clients
WHERE
    deleted_at IS NULL
    AND tenant_id = $1;

-- name: GetAuthClient :one
SELECT
    id,
    name,
    secret,
    last_modified_at,
    last_modified_by,
    tenant_id
FROM
    auth_clients
WHERE
    deleted_at IS NULL
    AND tenant_id = $1
    AND id = $2;

-- name: UpdateAuthClient :one
UPDATE
    auth_clients
SET
    name = $3,
    last_modified_at = $4,
    last_modified_by = $5
WHERE
    deleted_at IS NULL
    AND tenant_id = $1
    AND id = $2
RETURNING
    id,
    name,
    secret,
    last_modified_at,
    last_modified_by,
    tenant_id;

-- name: DeleteAuthClient :exec
UPDATE
    auth_clients
SET
    deleted_at = $3
WHERE
    deleted_at IS NULL
    AND tenant_id = $1
    AND id = $2;

