// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: food_item.sql

package db

import (
	"context"
)

const getFoodItem = `-- name: GetFoodItem :one
SELECT id, name, carbon_footprint, typology_id, source_id FROM food_item
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetFoodItem(ctx context.Context, id int32) (FoodItem, error) {
	row := q.db.QueryRow(ctx, getFoodItem, id)
	var i FoodItem
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CarbonFootprint,
		&i.TypologyID,
		&i.SourceID,
	)
	return i, err
}
