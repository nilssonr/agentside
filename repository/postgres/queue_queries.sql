-- name: InsertQueue :one
INSERT INTO queues (name, tenant_id, last_modified_at, last_modified_by)
    VALUES ($1, $2, $3, $4)
RETURNING
    id, name, tenant_id, last_modified_at, last_modified_by;

-- name: GetQueues :many
SELECT
    id,
    name,
    tenant_id,
    last_modified_at,
    last_modified_by
FROM
    queues
WHERE
    deleted_at IS NULL
    AND tenant_id = $1;

-- name: GetQueue :one
SELECT
    id,
    name,
    tenant_id,
    last_modified_at,
    last_modified_by
FROM
    queues
WHERE
    deleted_at IS NULL
    AND tenant_id = $1
    AND id = $2;

-- name: UpdateQueue :one
UPDATE
    queues
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
    tenant_id,
    last_modified_at,
    last_modified_by;

-- name: DeleteQueue :exec
UPDATE
    queues
SET
    deleted_at = $3
WHERE
    deleted_at IS NULL
    AND tenant_id = $1
    AND id = $2;

