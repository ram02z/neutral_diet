-- name: CreateFoodItem :one
INSERT INTO food_item (name, typology_id, suggested_cf)
    VALUES ($1, $2, $3)
RETURNING
    id;
