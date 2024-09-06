-- name: CreateSkill :one
INSERT INTO skills (name, tenant_id, last_modified_at, last_modified_by)
    VALUES ($1, $2, $3, $4)
RETURNING
    id, name, tenant_id, last_modified_at, last_modified_by;

-- name: GetSkills :many
SELECT
    id,
    name,
    tenant_id,
    last_modified_at,
    last_modified_by
FROM
    skills
WHERE
    deleted_at IS NULL
    AND tenant_id = $1;

-- name: GetSkill :one
SELECT
    id,
    name,
    tenant_id,
    last_modified_at,
    last_modified_by
FROM
    skills
WHERE
    deleted_at IS NULL
    AND tenant_id = $1
    AND id = $2;

-- name: UpdateSkill :one
UPDATE
    skills
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

-- name: DeleteSkill :exec
UPDATE
    skills
SET
    deleted_at = $3
WHERE
    delete_at IS NULL
    AND tenant_id = $1
    AND id = $2;

