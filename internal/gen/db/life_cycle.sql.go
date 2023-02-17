// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: life_cycle.sql

package db

import (
	"context"

	"github.com/shopspring/decimal"
)

const createLifeCycle = `-- name: CreateLifeCycle :one
INSERT INTO life_cycle (carbon_footprint, food_item_id, source_id)
    VALUES ($1, $2, $3)
RETURNING
    id
`

type CreateLifeCycleParams struct {
	CarbonFootprint decimal.Decimal
	FoodItemID      int32
	SourceID        int32
}

func (q *Queries) CreateLifeCycle(ctx context.Context, arg CreateLifeCycleParams) (int32, error) {
	row := q.db.QueryRow(ctx, createLifeCycle, arg.CarbonFootprint, arg.FoodItemID, arg.SourceID)
	var id int32
	err := row.Scan(&id)
	return id, err
}
