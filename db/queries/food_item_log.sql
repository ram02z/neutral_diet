-- name: AddFoodItemToLog :one
INSERT INTO "food_item_log" (food_item_id, weight, user_id, log_date, weight_unit, region, carbon_footprint)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING
    id;

-- name: UpdateFoodItemFromLog :exec
UPDATE
    "food_item_log"
SET
    weight = $3,
    weight_unit = $4,
    carbon_footprint = $5
WHERE
    user_id = $1
    AND id = $2;

-- name: GetFoodItemIdByLogId :one
SELECT
    food_item_id
FROM
    "food_item_log"
WHERE
    id = $1;

-- name: GetFoodItemLogByDate :many
SELECT
    l.id,
    f.name,
    l.food_item_id,
    l.region,
    l.weight,
    l.weight_unit,
    l.log_date,
    l.carbon_footprint
FROM
    food_item_log l
    INNER JOIN food_item f ON l.food_item_id = f.id
WHERE
    user_id = $1
    AND log_date = $2;

-- name: GetFoodItemLog :many
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
    user_id = $1;

-- name: GetDailyAverageCarbonFootprintByDietaryRequirement :one
SELECT
    average_carbon_footprint
FROM
    daily_user_average_by_dietary_requirement
WHERE
    dietary_requirement = $1
LIMIT 1;

-- name: GetDailyAverageCarbonFootprint :one
SELECT
    average_carbon_footprint
FROM
    daily_user_average
LIMIT 1;

-- name: GetFoodItemLogStreak :one
WITH dates AS (
    SELECT DISTINCT
        log_date as date
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
    CURRENT_DATE < MAX(date) + 1 AS active
FROM
    GROUPS
GROUP BY
    grp
HAVING
    COUNT(*) > 1
LIMIT 1;

-- name: DeleteFoodItemFromLog :exec
DELETE FROM "food_item_log"
WHERE user_id = $1
    AND id = $2;

-- name: DeleteUserLog :exec
DELETE FROM "food_item_log"
WHERE user_id = $1;
