// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: goal.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

const addCarbonFootprintGoal = `-- name: AddCarbonFootprintGoal :one
INSERT INTO "carbon_footprint_goal" (user_id, description, start_date, end_date, start_carbon_footprint, target_carbon_footprint)
    VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
    id
`

type AddCarbonFootprintGoalParams struct {
	UserID                int32
	Description           string
	StartDate             pgtype.Date
	EndDate               pgtype.Date
	StartCarbonFootprint  decimal.Decimal
	TargetCarbonFootprint decimal.Decimal
}

func (q *Queries) AddCarbonFootprintGoal(ctx context.Context, arg AddCarbonFootprintGoalParams) (int32, error) {
	row := q.db.QueryRow(ctx, addCarbonFootprintGoal,
		arg.UserID,
		arg.Description,
		arg.StartDate,
		arg.EndDate,
		arg.StartCarbonFootprint,
		arg.TargetCarbonFootprint,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteCarbonFootprintGoal = `-- name: DeleteCarbonFootprintGoal :exec
DELETE FROM "carbon_footprint_goal"
WHERE user_id = $1
    AND id = $2
`

type DeleteCarbonFootprintGoalParams struct {
	UserID int32
	ID     int32
}

func (q *Queries) DeleteCarbonFootprintGoal(ctx context.Context, arg DeleteCarbonFootprintGoalParams) error {
	_, err := q.db.Exec(ctx, deleteCarbonFootprintGoal, arg.UserID, arg.ID)
	return err
}

const getCarbonFootprintGoals = `-- name: GetCarbonFootprintGoals :many
SELECT
    id, user_id, description, start_date, end_date, start_carbon_footprint, target_carbon_footprint, completed
FROM
    "carbon_footprint_goal"
WHERE
    user_id = $1
ORDER BY
    end_date
`

func (q *Queries) GetCarbonFootprintGoals(ctx context.Context, userID int32) ([]CarbonFootprintGoal, error) {
	rows, err := q.db.Query(ctx, getCarbonFootprintGoals, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CarbonFootprintGoal
	for rows.Next() {
		var i CarbonFootprintGoal
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Description,
			&i.StartDate,
			&i.EndDate,
			&i.StartCarbonFootprint,
			&i.TargetCarbonFootprint,
			&i.Completed,
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

const updateCarbonFootprintGoal = `-- name: UpdateCarbonFootprintGoal :exec
UPDATE
    "carbon_footprint_goal"
SET
    completed = $3
WHERE
    user_id = $1
    AND id = $2
`

type UpdateCarbonFootprintGoalParams struct {
	UserID    int32
	ID        int32
	Completed bool
}

func (q *Queries) UpdateCarbonFootprintGoal(ctx context.Context, arg UpdateCarbonFootprintGoalParams) error {
	_, err := q.db.Exec(ctx, updateCarbonFootprintGoal, arg.UserID, arg.ID, arg.Completed)
	return err
}
