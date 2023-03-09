// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: food_item_log.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

const addFoodItemToLog = `-- name: AddFoodItemToLog :one
INSERT INTO "food_item_log" (food_item_id, weight, user_id, log_date, weight_unit, region, carbon_footprint, meal)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING
    id
`

type AddFoodItemToLogParams struct {
	FoodItemID      int32
	Weight          decimal.Decimal
	UserID          int32
	LogDate         pgtype.Date
	WeightUnit      int32
	Region          int32
	CarbonFootprint decimal.Decimal
	Meal            int32
}

func (q *Queries) AddFoodItemToLog(ctx context.Context, arg AddFoodItemToLogParams) (int32, error) {
	row := q.db.QueryRow(ctx, addFoodItemToLog,
		arg.FoodItemID,
		arg.Weight,
		arg.UserID,
		arg.LogDate,
		arg.WeightUnit,
		arg.Region,
		arg.CarbonFootprint,
		arg.Meal,
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

const getDailyAverageCarbonFootprint = `-- name: GetDailyAverageCarbonFootprint :one
SELECT
    average_carbon_footprint
FROM
    daily_user_average
LIMIT 1
`

func (q *Queries) GetDailyAverageCarbonFootprint(ctx context.Context) (decimal.Decimal, error) {
	row := q.db.QueryRow(ctx, getDailyAverageCarbonFootprint)
	var average_carbon_footprint decimal.Decimal
	err := row.Scan(&average_carbon_footprint)
	return average_carbon_footprint, err
}

const getDailyAverageCarbonFootprintByDietaryRequirement = `-- name: GetDailyAverageCarbonFootprintByDietaryRequirement :one
SELECT
    average_carbon_footprint
FROM
    daily_user_average_by_dietary_requirement
WHERE
    dietary_requirement = $1
LIMIT 1
`

func (q *Queries) GetDailyAverageCarbonFootprintByDietaryRequirement(ctx context.Context, dietaryRequirement int32) (decimal.Decimal, error) {
	row := q.db.QueryRow(ctx, getDailyAverageCarbonFootprintByDietaryRequirement, dietaryRequirement)
	var average_carbon_footprint decimal.Decimal
	err := row.Scan(&average_carbon_footprint)
	return average_carbon_footprint, err
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

const getFoodItemLog = `-- name: GetFoodItemLog :many
SELECT
    l.id,
    f.name,
    l.food_item_id,
    l.region,
    l.weight,
    l.weight_unit,
    l.carbon_footprint
FROM
    food_item_log l
    INNER JOIN food_item f ON l.food_item_id = f.id
WHERE
    user_id = $1
`

type GetFoodItemLogRow struct {
	ID              int32
	Name            string
	FoodItemID      int32
	Region          int32
	Weight          decimal.Decimal
	WeightUnit      int32
	CarbonFootprint decimal.Decimal
}

func (q *Queries) GetFoodItemLog(ctx context.Context, userID int32) ([]GetFoodItemLogRow, error) {
	rows, err := q.db.Query(ctx, getFoodItemLog, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFoodItemLogRow
	for rows.Next() {
		var i GetFoodItemLogRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.FoodItemID,
			&i.Region,
			&i.Weight,
			&i.WeightUnit,
			&i.CarbonFootprint,
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

const getFoodItemLogByDate = `-- name: GetFoodItemLogByDate :many
SELECT
    l.id,
    f.name,
    l.food_item_id,
    l.region,
    l.weight,
    l.weight_unit,
    l.log_date,
    l.meal,
    l.carbon_footprint
FROM
    food_item_log l
    INNER JOIN food_item f ON l.food_item_id = f.id
WHERE
    user_id = $1
    AND log_date = $2
`

type GetFoodItemLogByDateParams struct {
	UserID  int32
	LogDate pgtype.Date
}

type GetFoodItemLogByDateRow struct {
	ID              int32
	Name            string
	FoodItemID      int32
	Region          int32
	Weight          decimal.Decimal
	WeightUnit      int32
	LogDate         pgtype.Date
	Meal            int32
	CarbonFootprint decimal.Decimal
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
			&i.Name,
			&i.FoodItemID,
			&i.Region,
			&i.Weight,
			&i.WeightUnit,
			&i.LogDate,
			&i.Meal,
			&i.CarbonFootprint,
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

const getFoodItemLogStreak = `-- name: GetFoodItemLogStreak :one
WITH dates AS (
    SELECT DISTINCT
        log_date AS date
    FROM
        food_item_log
    WHERE
        user_id = $1
),
GROUPS AS (
    SELECT
        ROW_NUMBER() OVER (ORDER BY date) AS rn,
    date + - ROW_NUMBER() OVER (ORDER BY date) * INTERVAL '1 day' AS grp,
    date
FROM
    dates
)
SELECT
    COUNT(*) AS consecutive_dates,
    MIN(date)::date AS start_date,
    MAX(date)::date AS end_date,
    CURRENT_DATE <= MAX(date) + 1 AS active -- must be more than 1 to be a streak
FROM
    GROUPS
GROUP BY
    grp
HAVING
    COUNT(*) > 1
LIMIT 1
`

type GetFoodItemLogStreakRow struct {
	ConsecutiveDates int64
	StartDate        pgtype.Date
	EndDate          pgtype.Date
	Active           bool
}

// Generate "groups" of dates by subtracting the
// date's row number (no gaps) from the date itself
// (with potential gaps). Whenever there is a gap,
// there will be a new group
func (q *Queries) GetFoodItemLogStreak(ctx context.Context, userID int32) (GetFoodItemLogStreakRow, error) {
	row := q.db.QueryRow(ctx, getFoodItemLogStreak, userID)
	var i GetFoodItemLogStreakRow
	err := row.Scan(
		&i.ConsecutiveDates,
		&i.StartDate,
		&i.EndDate,
		&i.Active,
	)
	return i, err
}

const updateFoodItemFromLog = `-- name: UpdateFoodItemFromLog :exec
UPDATE
    "food_item_log"
SET
    weight = $3,
    weight_unit = $4,
    carbon_footprint = $5,
    meal = $6
WHERE
    user_id = $1
    AND id = $2
`

type UpdateFoodItemFromLogParams struct {
	UserID          int32
	ID              int32
	Weight          decimal.Decimal
	WeightUnit      int32
	CarbonFootprint decimal.Decimal
	Meal            int32
}

func (q *Queries) UpdateFoodItemFromLog(ctx context.Context, arg UpdateFoodItemFromLogParams) error {
	_, err := q.db.Exec(ctx, updateFoodItemFromLog,
		arg.UserID,
		arg.ID,
		arg.Weight,
		arg.WeightUnit,
		arg.CarbonFootprint,
		arg.Meal,
	)
	return err
}
