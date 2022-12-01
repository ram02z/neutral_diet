-- name: CreateLifeCycle :one
INSERT INTO life_cycle (carbon_footprint, food_item_id, source_id)
    VALUES ($1, $2, $3)
RETURNING
    id;
