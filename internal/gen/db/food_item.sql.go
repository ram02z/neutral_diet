// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: food_item.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
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

const getFoodItemInfo = `-- name: GetFoodItemInfo :one
SELECT
    t.name AS typology_name,
    st.name AS sub_typology_name,
    COUNT(l.food_item_id) AS n
FROM
    life_cycle l
    INNER JOIN food_item f ON l.food_item_id = f.id
    INNER JOIN typology t ON f.typology_id = t.id
    LEFT JOIN sub_typology st ON t.sub_typology_id = st.id
WHERE
    l.food_item_id = $1
GROUP BY
    t.name,
    st.name
`

type GetFoodItemInfoRow struct {
	TypologyName    string
	SubTypologyName pgtype.Text
	N               int64
}

func (q *Queries) GetFoodItemInfo(ctx context.Context, foodItemID int32) (GetFoodItemInfoRow, error) {
	row := q.db.QueryRow(ctx, getFoodItemInfo, foodItemID)
	var i GetFoodItemInfoRow
	err := row.Scan(&i.TypologyName, &i.SubTypologyName, &i.N)
	return i, err
}

const getFoodItemInfoByRegion = `-- name: GetFoodItemInfoByRegion :one
SELECT
    t.name AS typology_name,
    st.name AS sub_typology_name,
    COUNT(l.food_item_id) AS n
FROM
    life_cycle l
    INNER JOIN source s ON l.source_id = s.id
    INNER JOIN food_item f ON l.food_item_id = f.id
    INNER JOIN typology t ON f.typology_id = t.id
    LEFT JOIN sub_typology st ON t.sub_typology_id = st.id
WHERE
    l.food_item_id = $1
    AND s.region = $2
GROUP BY
    t.name,
    st.name
`

type GetFoodItemInfoByRegionParams struct {
	FoodItemID int32
	Region     int32
}

type GetFoodItemInfoByRegionRow struct {
	TypologyName    string
	SubTypologyName pgtype.Text
	N               int64
}

func (q *Queries) GetFoodItemInfoByRegion(ctx context.Context, arg GetFoodItemInfoByRegionParams) (GetFoodItemInfoByRegionRow, error) {
	row := q.db.QueryRow(ctx, getFoodItemInfoByRegion, arg.FoodItemID, arg.Region)
	var i GetFoodItemInfoByRegionRow
	err := row.Scan(&i.TypologyName, &i.SubTypologyName, &i.N)
	return i, err
}
