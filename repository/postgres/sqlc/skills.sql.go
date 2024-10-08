// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: skills.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteSkill = `-- name: DeleteSkill :exec
UPDATE
    skills
SET
    deleted_at = $3
WHERE
    delete_at IS NULL
    AND tenant_id = $1
    AND id = $2
`

type DeleteSkillParams struct {
	TenantID  string
	ID        string
	DeletedAt pgtype.Timestamptz
}

func (q *Queries) DeleteSkill(ctx context.Context, arg DeleteSkillParams) error {
	_, err := q.db.Exec(ctx, deleteSkill, arg.TenantID, arg.ID, arg.DeletedAt)
	return err
}

const getSkill = `-- name: GetSkill :one
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
    AND id = $2
`

type GetSkillParams struct {
	TenantID string
	ID       string
}

type GetSkillRow struct {
	ID             string
	Name           string
	TenantID       string
	LastModifiedAt pgtype.Timestamptz
	LastModifiedBy string
}

func (q *Queries) GetSkill(ctx context.Context, arg GetSkillParams) (GetSkillRow, error) {
	row := q.db.QueryRow(ctx, getSkill, arg.TenantID, arg.ID)
	var i GetSkillRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.TenantID,
		&i.LastModifiedAt,
		&i.LastModifiedBy,
	)
	return i, err
}

const getSkills = `-- name: GetSkills :many
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
`

type GetSkillsRow struct {
	ID             string
	Name           string
	TenantID       string
	LastModifiedAt pgtype.Timestamptz
	LastModifiedBy string
}

func (q *Queries) GetSkills(ctx context.Context, tenantID string) ([]GetSkillsRow, error) {
	rows, err := q.db.Query(ctx, getSkills, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetSkillsRow
	for rows.Next() {
		var i GetSkillsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.TenantID,
			&i.LastModifiedAt,
			&i.LastModifiedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertSkill = `-- name: InsertSkill :one
INSERT INTO skills(name, tenant_id, last_modified_at, last_modified_by)
    VALUES ($1, $2, $3, $4)
RETURNING
    id, name, tenant_id, last_modified_at, last_modified_by
`

type InsertSkillParams struct {
	Name           string
	TenantID       string
	LastModifiedAt pgtype.Timestamptz
	LastModifiedBy string
}

type InsertSkillRow struct {
	ID             string
	Name           string
	TenantID       string
	LastModifiedAt pgtype.Timestamptz
	LastModifiedBy string
}

func (q *Queries) InsertSkill(ctx context.Context, arg InsertSkillParams) (InsertSkillRow, error) {
	row := q.db.QueryRow(ctx, insertSkill,
		arg.Name,
		arg.TenantID,
		arg.LastModifiedAt,
		arg.LastModifiedBy,
	)
	var i InsertSkillRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.TenantID,
		&i.LastModifiedAt,
		&i.LastModifiedBy,
	)
	return i, err
}

const updateSkill = `-- name: UpdateSkill :one
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
    last_modified_by
`

type UpdateSkillParams struct {
	TenantID       string
	ID             string
	Name           string
	LastModifiedAt pgtype.Timestamptz
	LastModifiedBy string
}

type UpdateSkillRow struct {
	ID             string
	Name           string
	TenantID       string
	LastModifiedAt pgtype.Timestamptz
	LastModifiedBy string
}

func (q *Queries) UpdateSkill(ctx context.Context, arg UpdateSkillParams) (UpdateSkillRow, error) {
	row := q.db.QueryRow(ctx, updateSkill,
		arg.TenantID,
		arg.ID,
		arg.Name,
		arg.LastModifiedAt,
		arg.LastModifiedBy,
	)
	var i UpdateSkillRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.TenantID,
		&i.LastModifiedAt,
		&i.LastModifiedBy,
	)
	return i, err
}
