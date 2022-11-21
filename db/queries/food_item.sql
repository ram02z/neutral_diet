-- name: CreateFoodItem :one
INSERT INTO food_item (name, carbon_footprint, typology_id, source_id)
    VALUES ($1, $2, $3, $4)
RETURNING
    id;
