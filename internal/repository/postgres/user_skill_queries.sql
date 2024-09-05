-- name: UpsertUserSkill :one
WITH upserted AS (
INSERT INTO user_skills (user_id, skill_id, skill_level)
        VALUES ($1, $2, $3)
    ON CONFLICT (user_id, skill_id)
        DO UPDATE SET
            skill_level = EXCLUDED.skill_level
        RETURNING
            skill_id, skill_level)
        SELECT
            u.skill_id AS id,
            s.name AS name,
            u.skill_level AS skill_level
        FROM
            upserted u
            JOIN skills s ON s.id = u.skill_id;

-- name: GetUserSkills :many
SELECT
    s.id AS id,
    s.name AS name,
    us.skill_level AS skill_level
FROM
    user_skills us
    JOIN skills s ON s.id = us.skill_id
WHERE
    us.user_id = $1;

-- name: GetUserSkill :one
SELECT
    s.id AS id,
    s.name AS name,
    us.skill_level AS skill_level
FROM
    user_skills us
    JOIN skills s ON s.id = us.skill_id
WHERE
    us.user_id = $1
    AND s.id = $2
LIMIT 1;

-- name: DeleteUserSkill :exec
DELETE FROM user_skills
WHERE user_id = $1
    AND skill_id = $2;

