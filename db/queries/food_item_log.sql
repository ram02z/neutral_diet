-- name: AddFoodItemToLog :one
INSERT INTO "food_item_log" (food_item_id, weight, carbon_footprint, user_id, log_date)
    VALUES ($1, $2, $3, $4, $5)
RETURNING
    id;

-- name: GetFoodItemLogByDate :many
SELECT
    *
FROM
    "food_item_log"
WHERE
    log_date = $1;

-- name: DeleteFoodItemFromLog :exec
DELETE FROM "food_item_log"
WHERE user_id = $1
    AND id = $2;

-- name: DeleteUserLog :exec
DELETE FROM "food_item_log"
WHERE user_id = $1;
