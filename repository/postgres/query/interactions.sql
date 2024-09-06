-- name: InsertInteraction :one
INSERT INTO interactions (type, queue_id, state, state_modified_at, user_id, tenant_id, created_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING
    id, type, queue_id, state, state_modified_at, user_id, tenant_id, created_at;

-- name: GetInteractions :many
SELECT
    id,
    type,
    queue_id,
    state,
    state_modified_at,
    user_id,
    tenant_id,
    created_at
FROM
    interactions
WHERE
    tenant_id = $1;

-- name: GetInteraction :one
SELECT
    id,
    type,
    queue_id,
    state,
    state_modified_at,
    user_id,
    tenant_id,
    created_at
FROM
    interactions
WHERE
    tenant_id = $1
    AND id = $2;
