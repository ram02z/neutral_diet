-- name: AddFoodItemToLog :one
INSERT INTO "food_item_log" (food_item_id, quantity, user_id, log_date, unit, region, carbon_footprint, meal)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING
    id;

-- name: UpdateFoodItemFromLog :exec
UPDATE
    "food_item_log"
SET
    quantity = $3,
    unit = $4,
    carbon_footprint = $5,
    meal = $6
WHERE
    user_id = $1
    AND id = $2;

-- name: GetDailyCarbonFootprint :many
SELECT
    sum(carbon_footprint)::decimal AS carbon_footprint,
    log_date,
    meal
FROM
    food_item_log
WHERE
    user_id = $1
GROUP BY
    log_date,
    meal
ORDER BY
    log_date;

-- name: GetLoggedDaysInMonth :many
SELECT DISTINCT
    DATE_PART('day', log_date)::int AS day
FROM
    food_item_log
WHERE
    DATE_PART('month', log_date) = @month::int
    AND DATE_PART('year', log_date) = @year::int
    AND user_id = $1;

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
    l.quantity,
    l.unit,
    l.log_date,
    l.meal,
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
    l.quantity,
    l.unit,
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
        log_date AS date
    FROM
        food_item_log
    WHERE
        user_id = $1
),
-- Generate "groups" of dates by subtracting the
-- date's row number (no gaps) from the date itself
-- (with potential gaps). Whenever there is a gap,
-- there will be a new group
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
LIMIT 1;

-- name: DeleteFoodItemFromLog :exec
DELETE FROM "food_item_log"
WHERE user_id = $1
    AND id = $2;
