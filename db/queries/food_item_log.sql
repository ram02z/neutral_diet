-- name: AddFoodItemToLog :one
INSERT INTO "food_item_log" (food_item_id, weight, carbon_footprint, user_id, log_date)
    VALUES ($1, $2, $3, $4, $5)
RETURNING
    id;

-- name: UpdateFoodItemFromLog :exec
UPDATE
    "food_item_log"
SET
    weight = $3,
    carbon_footprint = $4
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
    l.food_item_id,
    f.name,
    l.weight,
    l.carbon_footprint,
    l.log_date
FROM
    food_item_log l
    INNER JOIN food_item f ON l.food_item_id = f.id
WHERE
    user_id = $1
    AND log_date = $2;

-- name: DeleteFoodItemFromLog :exec
DELETE FROM "food_item_log"
WHERE user_id = $1
    AND id = $2;

-- name: DeleteUserLog :exec
DELETE FROM "food_item_log"
WHERE user_id = $1;
