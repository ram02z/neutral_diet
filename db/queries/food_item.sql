-- name: CreateFoodItem :one
INSERT INTO food_item (name, carbon_footprint, typology_id, source_id)
    VALUES ($1, $2, $3, $4)
RETURNING
    id;

-- name: ListFoodItems :many
SELECT
    id,
    name,
    carbon_footprint
FROM
    food_item
ORDER BY
    ID;
