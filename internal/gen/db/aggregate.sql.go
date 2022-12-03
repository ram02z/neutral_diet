// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: aggregate.sql

package db

import (
	"context"

	"github.com/jackc/pgtype"
)

const listAggregateFoodItems = `-- name: ListAggregateFoodItems :many
SELECT
    f.id AS food_item_id,
    COUNT(*) AS n,
    CAST(PERCENTILE_CONT(0.50) WITHIN GROUP (ORDER BY l.carbon_footprint) AS DECIMAL) AS median_carbon_footprint
FROM
    life_cycle l
    INNER JOIN food_item f ON l.food_item_id = f.id
    INNER JOIN source s ON l.source_id = s.id
WHERE
    s.region_name = $1
GROUP BY
    f.id
`

type ListAggregateFoodItemsRow struct {
	FoodItemID            int32
	N                     int64
	MedianCarbonFootprint pgtype.Numeric
}

func (q *Queries) ListAggregateFoodItems(ctx context.Context, regionName string) ([]ListAggregateFoodItemsRow, error) {
	rows, err := q.db.Query(ctx, listAggregateFoodItems, regionName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAggregateFoodItemsRow
	for rows.Next() {
		var i ListAggregateFoodItemsRow
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

const listAggregateSubTypologies = `-- name: ListAggregateSubTypologies :many
SELECT
    st.id AS sub_typology_id,
    COUNT(*) AS n,
    CAST(PERCENTILE_CONT(0.50) WITHIN GROUP (ORDER BY l.carbon_footprint) AS DECIMAL) AS median_carbon_footprint
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

type ListAggregateSubTypologiesRow struct {
	SubTypologyID         int32
	N                     int64
	MedianCarbonFootprint pgtype.Numeric
}

func (q *Queries) ListAggregateSubTypologies(ctx context.Context, regionName string) ([]ListAggregateSubTypologiesRow, error) {
	rows, err := q.db.Query(ctx, listAggregateSubTypologies, regionName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAggregateSubTypologiesRow
	for rows.Next() {
		var i ListAggregateSubTypologiesRow
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

const listAggregateTypologies = `-- name: ListAggregateTypologies :many
SELECT
    t.id AS typology_id,
    COUNT(*) AS n,
    CAST(PERCENTILE_CONT(0.50) WITHIN GROUP (ORDER BY l.carbon_footprint) AS DECIMAL) AS median_carbon_footprint
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

type ListAggregateTypologiesRow struct {
	TypologyID            int32
	N                     int64
	MedianCarbonFootprint pgtype.Numeric
}

func (q *Queries) ListAggregateTypologies(ctx context.Context, regionName string) ([]ListAggregateTypologiesRow, error) {
	rows, err := q.db.Query(ctx, listAggregateTypologies, regionName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAggregateTypologiesRow
	for rows.Next() {
		var i ListAggregateTypologiesRow
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
