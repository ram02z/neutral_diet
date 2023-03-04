-- name: CreateTypology :one
INSERT INTO typology (name, sub_typology_id)
    VALUES ($1, $2)
RETURNING
    id;

-- name: ListTypologies :many
SELECT * FROM typology;
