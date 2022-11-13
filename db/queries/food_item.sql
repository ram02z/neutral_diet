-- name: GetFoodItem :one
SELECT * FROM food_item
WHERE id = $1 LIMIT 1;
