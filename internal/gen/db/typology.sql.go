// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: typology.sql

package db

import (
	"context"
	"database/sql"
)

const createTypology = `-- name: CreateTypology :one
INSERT INTO typology (name, sub_typology_id)
    VALUES ($1, $2)
RETURNING
    id
`

type CreateTypologyParams struct {
	Name          string
	SubTypologyID sql.NullInt32
}

func (q *Queries) CreateTypology(ctx context.Context, arg CreateTypologyParams) (int32, error) {
	row := q.db.QueryRow(ctx, createTypology, arg.Name, arg.SubTypologyID)
	var id int32
	err := row.Scan(&id)
	return id, err
}
