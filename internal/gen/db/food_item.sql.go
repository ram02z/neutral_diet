// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: food_item.sql

package db

import (
	"context"

	"github.com/jackc/pgtype"
)

const createFoodItem = `-- name: CreateFoodItem :one
INSERT INTO food_item (name, carbon_footprint, typology_id, source_id)
    VALUES ($1, $2, $3, $4)
RETURNING
    id
`

type CreateFoodItemParams struct {
	Name            string
	CarbonFootprint pgtype.Numeric
	TypologyID      int32
	SourceID        int32
}

func (q *Queries) CreateFoodItem(ctx context.Context, arg CreateFoodItemParams) (int32, error) {
	row := q.db.QueryRow(ctx, createFoodItem,
		arg.Name,
		arg.CarbonFootprint,
		arg.TypologyID,
		arg.SourceID,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const listAggregateFoodItems = `-- name: ListAggregateFoodItems :many
SELECT
    name, n, median_carbon_footprint, typology_id, id
FROM
    aggregate_food_item
`

func (q *Queries) ListAggregateFoodItems(ctx context.Context) ([]AggregateFoodItem, error) {
	rows, err := q.db.Query(ctx, listAggregateFoodItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AggregateFoodItem
	for rows.Next() {
		var i AggregateFoodItem
		if err := rows.Scan(
			&i.Name,
			&i.N,
			&i.MedianCarbonFootprint,
			&i.TypologyID,
			&i.ID,
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

const listFoodItems = `-- name: ListFoodItems :many
SELECT
    id,
    name,
    carbon_footprint
FROM
    food_item
ORDER BY
    ID
`

type ListFoodItemsRow struct {
	ID              int32
	Name            string
	CarbonFootprint pgtype.Numeric
}

func (q *Queries) ListFoodItems(ctx context.Context) ([]ListFoodItemsRow, error) {
	rows, err := q.db.Query(ctx, listFoodItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListFoodItemsRow
	for rows.Next() {
		var i ListFoodItemsRow
		if err := rows.Scan(&i.ID, &i.Name, &i.CarbonFootprint); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
