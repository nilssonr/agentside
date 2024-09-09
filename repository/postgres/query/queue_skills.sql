-- name: UpsertQueueSkill :one
WITH upserted AS (
INSERT INTO queue_skills(queue_id, skill_id, level, choice)
        VALUES ($1, $2, $3, $4)
    ON CONFLICT (queue_id, skill_id)
        DO UPDATE SET level = EXCLUDED.level, choice = EXCLUDED.choice
    RETURNING
        skill_id, level, choice)
        SELECT
            u.skill_id AS id,
            s.name AS name,
            u.level AS level,
            u.choice AS choice
        FROM
            upserted u
            JOIN skills s ON s.id = u.skill_id;

-- name: GetQueueSkills :many
SELECT
    s.id,
    s.name,
    qs.level,
    qs.choice
FROM
    queue_skills qs
    JOIN skills s ON s.id = qs.skill_id
WHERE
    qs.queue_id = $1;

-- name: GetQueueSkill :one
SELECT
    s.id,
    s.name,
    qs.level,
    qs.choice
FROM
    queue_skills qs
    JOIN skills s ON s.id = qs.skill_id
WHERE
    qs.queue_id = $1
    AND s.id = $2;

-- name: DeleteQueueSkill :exec
DELETE FROM queue_skills
WHERE queue_id = $1
    AND skill_id = $2;

