-- name: CreateSource :one
INSERT INTO source (reference, year, region_name)
    VALUES ($1, $2, $3)
RETURNING
    id;
