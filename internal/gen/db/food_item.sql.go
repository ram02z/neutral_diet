// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: food_item.sql

package db

import (
	"context"
)

const createFoodItem = `-- name: CreateFoodItem :one
INSERT INTO food_item (name, typology_id, suggested_cf)
    VALUES ($1, $2, $3)
RETURNING
    id
`

type CreateFoodItemParams struct {
	Name        string
	TypologyID  int32
	SuggestedCf CfTypes
}

func (q *Queries) CreateFoodItem(ctx context.Context, arg CreateFoodItemParams) (int32, error) {
	row := q.db.QueryRow(ctx, createFoodItem, arg.Name, arg.TypologyID, arg.SuggestedCf)
	var id int32
	err := row.Scan(&id)
	return id, err
}
