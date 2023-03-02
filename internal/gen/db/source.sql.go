// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: source.sql

package db

import (
	"context"
)

const createSource = `-- name: CreateSource :one
INSERT INTO source (reference, year, region)
    VALUES ($1, $2, $3)
RETURNING
    id
`

type CreateSourceParams struct {
	Reference string
	Year      int32
	Region    int32
}

func (q *Queries) CreateSource(ctx context.Context, arg CreateSourceParams) (int32, error) {
	row := q.db.QueryRow(ctx, createSource, arg.Reference, arg.Year, arg.Region)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const listSourcesByFoodItem = `-- name: ListSourcesByFoodItem :many
SELECT
    s.reference,
    s.year,
    s.region
FROM
    life_cycle l
    INNER JOIN source s ON l.source_id = s.id
WHERE
    l.food_item_id = $1
GROUP BY
    s.id
`

type ListSourcesByFoodItemRow struct {
	Reference string
	Year      int32
	Region    int32
}

func (q *Queries) ListSourcesByFoodItem(ctx context.Context, foodItemID int32) ([]ListSourcesByFoodItemRow, error) {
	rows, err := q.db.Query(ctx, listSourcesByFoodItem, foodItemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListSourcesByFoodItemRow
	for rows.Next() {
		var i ListSourcesByFoodItemRow
		if err := rows.Scan(&i.Reference, &i.Year, &i.Region); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSourcesByFoodItemAndRegion = `-- name: ListSourcesByFoodItemAndRegion :many
SELECT
    s.reference,
    s.year,
    s.region
FROM
    life_cycle l
    INNER JOIN source s ON l.source_id = s.id
WHERE
    l.food_item_id = $1
    AND s.region = $2
GROUP BY
    s.id
`

type ListSourcesByFoodItemAndRegionParams struct {
	FoodItemID int32
	Region     int32
}

type ListSourcesByFoodItemAndRegionRow struct {
	Reference string
	Year      int32
	Region    int32
}

func (q *Queries) ListSourcesByFoodItemAndRegion(ctx context.Context, arg ListSourcesByFoodItemAndRegionParams) ([]ListSourcesByFoodItemAndRegionRow, error) {
	rows, err := q.db.Query(ctx, listSourcesByFoodItemAndRegion, arg.FoodItemID, arg.Region)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListSourcesByFoodItemAndRegionRow
	for rows.Next() {
		var i ListSourcesByFoodItemAndRegionRow
		if err := rows.Scan(&i.Reference, &i.Year, &i.Region); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
