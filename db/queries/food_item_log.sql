-- name: AddFoodItemToLog :one
INSERT INTO "food_item_log" (food_item_id, weight, carbon_footprint, user_id)
    VALUES ($1, $2, $3, $4)
RETURNING
    id;
