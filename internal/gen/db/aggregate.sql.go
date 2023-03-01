// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: aggregate.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

const getAggregateFoodItem = `-- name: GetAggregateFoodItem :one
SELECT
    food_item_id, n, median_carbon_footprint
FROM
    aggregate_food_item
WHERE
    food_item_id = $1
`

func (q *Queries) GetAggregateFoodItem(ctx context.Context, foodItemID int32) (AggregateFoodItem, error) {
	row := q.db.QueryRow(ctx, getAggregateFoodItem, foodItemID)
	var i AggregateFoodItem
	err := row.Scan(&i.FoodItemID, &i.N, &i.MedianCarbonFootprint)
	return i, err
}

const getRegionalAggregateFoodItem = `-- name: GetRegionalAggregateFoodItem :one
SELECT
    food_item_id, region, n, median_carbon_footprint
FROM
    regional_aggregate_food_item
WHERE
    food_item_id = $1
    AND region = $2
`

type GetRegionalAggregateFoodItemParams struct {
	FoodItemID int32
	Region     int32
}

func (q *Queries) GetRegionalAggregateFoodItem(ctx context.Context, arg GetRegionalAggregateFoodItemParams) (RegionalAggregateFoodItem, error) {
	row := q.db.QueryRow(ctx, getRegionalAggregateFoodItem, arg.FoodItemID, arg.Region)
	var i RegionalAggregateFoodItem
	err := row.Scan(
		&i.FoodItemID,
		&i.Region,
		&i.N,
		&i.MedianCarbonFootprint,
	)
	return i, err
}

const listAggregateFoodItems = `-- name: ListAggregateFoodItems :many
SELECT
    a.food_item_id AS id,
    f.name AS food_name,
    t.name AS typology_name,
    s.name AS sub_typology_name,
    a.n,
    ROUND(a.median_carbon_footprint, 3)::decimal AS median_carbon_footprint
FROM
    aggregate_food_item a
    INNER JOIN food_item f ON a.food_item_id = f.id
    INNER JOIN typology t ON f.typology_id = t.id
    LEFT JOIN sub_typology s ON t.sub_typology_id = s.id
`

type ListAggregateFoodItemsRow struct {
	ID                    int32
	FoodName              string
	TypologyName          string
	SubTypologyName       pgtype.Text
	N                     int64
	MedianCarbonFootprint decimal.Decimal
}

func (q *Queries) ListAggregateFoodItems(ctx context.Context) ([]ListAggregateFoodItemsRow, error) {
	rows, err := q.db.Query(ctx, listAggregateFoodItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAggregateFoodItemsRow
	for rows.Next() {
		var i ListAggregateFoodItemsRow
		if err := rows.Scan(
			&i.ID,
			&i.FoodName,
			&i.TypologyName,
			&i.SubTypologyName,
			&i.N,
			&i.MedianCarbonFootprint,
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

const listAggregateFoodItemsByRegion = `-- name: ListAggregateFoodItemsByRegion :many
SELECT
    a.food_item_id AS id,
    f.name AS food_name,
    t.name AS typology_name,
    s.name AS sub_typology_name,
    a.n,
    ROUND(a.median_carbon_footprint, 3)::decimal AS median_carbon_footprint
FROM
    regional_aggregate_food_item a
    INNER JOIN food_item f ON a.food_item_id = f.id
    INNER JOIN typology t ON f.typology_id = t.id
    LEFT JOIN sub_typology s ON t.sub_typology_id = s.id
WHERE
    a.region = $1
`

type ListAggregateFoodItemsByRegionRow struct {
	ID                    int32
	FoodName              string
	TypologyName          string
	SubTypologyName       pgtype.Text
	N                     int64
	MedianCarbonFootprint decimal.Decimal
}

func (q *Queries) ListAggregateFoodItemsByRegion(ctx context.Context, region int32) ([]ListAggregateFoodItemsByRegionRow, error) {
	rows, err := q.db.Query(ctx, listAggregateFoodItemsByRegion, region)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAggregateFoodItemsByRegionRow
	for rows.Next() {
		var i ListAggregateFoodItemsByRegionRow
		if err := rows.Scan(
			&i.ID,
			&i.FoodName,
			&i.TypologyName,
			&i.SubTypologyName,
			&i.N,
			&i.MedianCarbonFootprint,
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

const listAggregateSubTypologiesByRegion = `-- name: ListAggregateSubTypologiesByRegion :many
SELECT
    st.id AS sub_typology_id,
    COUNT(*) AS n,
    ROUND(PERCENTILE_CONT(0.50) WITHIN GROUP (ORDER BY l.carbon_footprint), 3)::decimal AS median_carbon_footprint
FROM
    life_cycle l
    INNER JOIN food_item f ON l.food_item_id = f.id
    INNER JOIN source s ON l.source_id = s.id
    INNER JOIN typology t ON f.typology_id = t.id
    INNER JOIN sub_typology st ON t.sub_typology_id = st.id
WHERE
    s.region = $1
GROUP BY
    st.id
`

type ListAggregateSubTypologiesByRegionRow struct {
	SubTypologyID         int32
	N                     int64
	MedianCarbonFootprint decimal.Decimal
}

func (q *Queries) ListAggregateSubTypologiesByRegion(ctx context.Context, region int32) ([]ListAggregateSubTypologiesByRegionRow, error) {
	rows, err := q.db.Query(ctx, listAggregateSubTypologiesByRegion, region)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAggregateSubTypologiesByRegionRow
	for rows.Next() {
		var i ListAggregateSubTypologiesByRegionRow
		if err := rows.Scan(&i.SubTypologyID, &i.N, &i.MedianCarbonFootprint); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAggregateTypologiesByRegion = `-- name: ListAggregateTypologiesByRegion :many
SELECT
    t.id AS typology_id,
    COUNT(*) AS n,
    ROUND(PERCENTILE_CONT(0.50) WITHIN GROUP (ORDER BY l.carbon_footprint), 3)::decimal AS median_carbon_footprint
FROM
    life_cycle l
    INNER JOIN food_item f ON l.food_item_id = f.id
    INNER JOIN source s ON l.source_id = s.id
    INNER JOIN typology t ON f.typology_id = t.id
WHERE
    s.region = $1
GROUP BY
    t.id
`

type ListAggregateTypologiesByRegionRow struct {
	TypologyID            int32
	N                     int64
	MedianCarbonFootprint decimal.Decimal
}

func (q *Queries) ListAggregateTypologiesByRegion(ctx context.Context, region int32) ([]ListAggregateTypologiesByRegionRow, error) {
	rows, err := q.db.Query(ctx, listAggregateTypologiesByRegion, region)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAggregateTypologiesByRegionRow
	for rows.Next() {
		var i ListAggregateTypologiesByRegionRow
		if err := rows.Scan(&i.TypologyID, &i.N, &i.MedianCarbonFootprint); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
