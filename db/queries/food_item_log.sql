-- name: AddFoodItemToLog :one
INSERT INTO "food_item_log" (food_item_id, weight, user_id, log_date, weight_unit, region)
    VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
    id;

-- name: UpdateFoodItemFromLog :exec
UPDATE
    "food_item_log"
SET
    weight = $3,
    weight_unit = $4
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
