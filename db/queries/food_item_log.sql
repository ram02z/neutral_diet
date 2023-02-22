-- name: AddFoodItemToLog :one
INSERT INTO "food_item_log" (food_item_id, weight, user_id, log_date, weight_unit)
    VALUES ($1, $2, $3, $4, $5)
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
    AND log_date = $2;

-- name: DeleteFoodItemFromLog :exec
DELETE FROM "food_item_log"
WHERE user_id = $1
    AND id = $2;

-- name: DeleteUserLog :exec
DELETE FROM "food_item_log"
WHERE user_id = $1;
