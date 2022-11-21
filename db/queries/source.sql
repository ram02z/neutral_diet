-- name: CreateSource :one
INSERT INTO source (reference, year, location)
    VALUES ($1, $2, $3)
RETURNING
    id;
