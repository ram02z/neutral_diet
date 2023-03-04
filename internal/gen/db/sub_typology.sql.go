// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: sub_typology.sql

package db

import (
	"context"
)

const createSubTypology = `-- name: CreateSubTypology :one
INSERT INTO sub_typology (name)
    VALUES ($1)
RETURNING
    id
`

func (q *Queries) CreateSubTypology(ctx context.Context, name string) (int32, error) {
	row := q.db.QueryRow(ctx, createSubTypology, name)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const listSubTypologyNames = `-- name: ListSubTypologyNames :many
SELECT DISTINCT(name) FROM sub_typology ORDER BY name ASC
`

func (q *Queries) ListSubTypologyNames(ctx context.Context) ([]string, error) {
	rows, err := q.db.Query(ctx, listSubTypologyNames)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
