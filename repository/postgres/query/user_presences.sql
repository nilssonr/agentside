-- name: UpsertUserPresence :one
INSERT INTO user_presences (user_id, presence, last_modified_at, last_modified_by)
    VALUES ($1, $2, $3, $4)
ON CONFLICT (user_id)
    DO UPDATE SET
        resence = EXCLUDED.presence, last_modified_at = EXCLUDED.last_modified_at, last_modified_by = EXCLUDED.last_modified_by
    RETURNING
        user_id, presence, last_modified_at, last_modified_by;

-- name: GetUserPresence :one
SELECT
    user_id,
    presence,
    last_modified_at,
    last_modified_by
FROM
    user_presences
WHERE
    user_id = $1;

-- name: DeleteUserPresence :exec
DELETE FROM user_presences
WHERE user_id = $1;

