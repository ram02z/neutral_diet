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

const getAggregateFoodItemById = `-- name: GetAggregateFoodItemById :one
SELECT
    food_item_id, n, median_carbon_footprint
FROM
    aggregate_food_item
WHERE
    food_item_id = $1
`

func (q *Queries) GetAggregateFoodItemById(ctx context.Context, foodItemID int32) (AggregateFoodItem, error) {
	row := q.db.QueryRow(ctx, getAggregateFoodItemById, foodItemID)
	var i AggregateFoodItem
	err := row.Scan(&i.FoodItemID, &i.N, &i.MedianCarbonFootprint)
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
    f.id AS food_item_id,
    COUNT(*) AS n,
    ROUND(PERCENTILE_CONT(0.50) WITHIN GROUP (ORDER BY l.carbon_footprint), 3)::decimal AS median_carbon_footprint
FROM
    life_cycle l
    INNER JOIN food_item f ON l.food_item_id = f.id
    INNER JOIN source s ON l.source_id = s.id
WHERE
    s.region_name = $1
GROUP BY
    f.id
`

type ListAggregateFoodItemsByRegionRow struct {
	FoodItemID            int32
	N                     int64
	MedianCarbonFootprint decimal.Decimal
}

func (q *Queries) ListAggregateFoodItemsByRegion(ctx context.Context, regionName string) ([]ListAggregateFoodItemsByRegionRow, error) {
	rows, err := q.db.Query(ctx, listAggregateFoodItemsByRegion, regionName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAggregateFoodItemsByRegionRow
	for rows.Next() {
		var i ListAggregateFoodItemsByRegionRow
		if err := rows.Scan(&i.FoodItemID, &i.N, &i.MedianCarbonFootprint); err != nil {
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
    s.region_name = $1
GROUP BY
    st.id
`

type ListAggregateSubTypologiesByRegionRow struct {
	SubTypologyID         int32
	N                     int64
	MedianCarbonFootprint decimal.Decimal
}

func (q *Queries) ListAggregateSubTypologiesByRegion(ctx context.Context, regionName string) ([]ListAggregateSubTypologiesByRegionRow, error) {
	rows, err := q.db.Query(ctx, listAggregateSubTypologiesByRegion, regionName)
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
    s.region_name = $1
GROUP BY
    t.id
`

type ListAggregateTypologiesByRegionRow struct {
	TypologyID            int32
	N                     int64
	MedianCarbonFootprint decimal.Decimal
}

func (q *Queries) ListAggregateTypologiesByRegion(ctx context.Context, regionName string) ([]ListAggregateTypologiesByRegionRow, error) {
	rows, err := q.db.Query(ctx, listAggregateTypologiesByRegion, regionName)
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
