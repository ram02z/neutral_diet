// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: food_item_log.sql

package db

import (
	"context"

	"github.com/shopspring/decimal"
)

const addFoodItemToLog = `-- name: AddFoodItemToLog :one
INSERT INTO "food_item_log" (food_item_id, weight, carbon_footprint, user_id)
    VALUES ($1, $2, $3, $4)
RETURNING
    id
`

type AddFoodItemToLogParams struct {
	FoodItemID      int32
	Weight          decimal.Decimal
	CarbonFootprint decimal.Decimal
	UserID          int32
}

func (q *Queries) AddFoodItemToLog(ctx context.Context, arg AddFoodItemToLogParams) (int32, error) {
	row := q.db.QueryRow(ctx, addFoodItemToLog,
		arg.FoodItemID,
		arg.Weight,
		arg.CarbonFootprint,
		arg.UserID,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteUserLog = `-- name: DeleteUserLog :exec
DELETE FROM "food_item_log"
WHERE user_id = $1
`

func (q *Queries) DeleteUserLog(ctx context.Context, userID int32) error {
	_, err := q.db.Exec(ctx, deleteUserLog, userID)
	return err
}
