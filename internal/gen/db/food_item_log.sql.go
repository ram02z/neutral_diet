// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: food_item_log.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

const addFoodItemToLog = `-- name: AddFoodItemToLog :one
INSERT INTO "food_item_log" (food_item_id, weight, user_id, log_date, weight_unit)
    VALUES ($1, $2, $3, $4, $5)
RETURNING
    id
`

type AddFoodItemToLogParams struct {
	FoodItemID int32
	Weight     decimal.Decimal
	UserID     int32
	LogDate    pgtype.Date
	WeightUnit WeightUnit
}

func (q *Queries) AddFoodItemToLog(ctx context.Context, arg AddFoodItemToLogParams) (int32, error) {
	row := q.db.QueryRow(ctx, addFoodItemToLog,
		arg.FoodItemID,
		arg.Weight,
		arg.UserID,
		arg.LogDate,
		arg.WeightUnit,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteFoodItemFromLog = `-- name: DeleteFoodItemFromLog :exec
DELETE FROM "food_item_log"
WHERE user_id = $1
    AND id = $2
`

type DeleteFoodItemFromLogParams struct {
	UserID int32
	ID     int32
}

func (q *Queries) DeleteFoodItemFromLog(ctx context.Context, arg DeleteFoodItemFromLogParams) error {
	_, err := q.db.Exec(ctx, deleteFoodItemFromLog, arg.UserID, arg.ID)
	return err
}

const deleteUserLog = `-- name: DeleteUserLog :exec
DELETE FROM "food_item_log"
WHERE user_id = $1
`

func (q *Queries) DeleteUserLog(ctx context.Context, userID int32) error {
	_, err := q.db.Exec(ctx, deleteUserLog, userID)
	return err
}

const getFoodItemIdByLogId = `-- name: GetFoodItemIdByLogId :one
SELECT
    food_item_id
FROM
    "food_item_log"
WHERE
    id = $1
`

func (q *Queries) GetFoodItemIdByLogId(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, getFoodItemIdByLogId, id)
	var food_item_id int32
	err := row.Scan(&food_item_id)
	return food_item_id, err
}

const getFoodItemLogByDate = `-- name: GetFoodItemLogByDate :many
SELECT
    l.id,
    l.food_item_id,
    f.name,
    l.weight,
    l.weight_unit,
    a.median_carbon_footprint,
    a.n,
    t.name AS typology_name,
    s.name AS sub_typology_name,
    l.log_date
FROM
    food_item_log l
    INNER JOIN food_item f ON l.food_item_id = f.id
    INNER JOIN aggregate_food_item a ON l.food_item_id = a.food_item_id
    INNER JOIN typology t ON f.typology_id = t.id
    LEFT JOIN sub_typology s ON t.sub_typology_id = s.id
WHERE
    user_id = $1
    AND log_date = $2
`

type GetFoodItemLogByDateParams struct {
	UserID  int32
	LogDate pgtype.Date
}

type GetFoodItemLogByDateRow struct {
	ID                    int32
	FoodItemID            int32
	Name                  string
	Weight                decimal.Decimal
	WeightUnit            WeightUnit
	MedianCarbonFootprint decimal.Decimal
	N                     int64
	TypologyName          string
	SubTypologyName       pgtype.Text
	LogDate               pgtype.Date
}

func (q *Queries) GetFoodItemLogByDate(ctx context.Context, arg GetFoodItemLogByDateParams) ([]GetFoodItemLogByDateRow, error) {
	rows, err := q.db.Query(ctx, getFoodItemLogByDate, arg.UserID, arg.LogDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFoodItemLogByDateRow
	for rows.Next() {
		var i GetFoodItemLogByDateRow
		if err := rows.Scan(
			&i.ID,
			&i.FoodItemID,
			&i.Name,
			&i.Weight,
			&i.WeightUnit,
			&i.MedianCarbonFootprint,
			&i.N,
			&i.TypologyName,
			&i.SubTypologyName,
			&i.LogDate,
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

const updateFoodItemFromLog = `-- name: UpdateFoodItemFromLog :exec
UPDATE
    "food_item_log"
SET
    weight = $3,
    weight_unit = $4
WHERE
    user_id = $1
    AND id = $2
`

type UpdateFoodItemFromLogParams struct {
	UserID     int32
	ID         int32
	Weight     decimal.Decimal
	WeightUnit WeightUnit
}

func (q *Queries) UpdateFoodItemFromLog(ctx context.Context, arg UpdateFoodItemFromLogParams) error {
	_, err := q.db.Exec(ctx, updateFoodItemFromLog,
		arg.UserID,
		arg.ID,
		arg.Weight,
		arg.WeightUnit,
	)
	return err
}
