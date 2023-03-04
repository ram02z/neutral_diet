-- name: CreateSubTypology :one
INSERT INTO sub_typology (name)
    VALUES ($1)
RETURNING
    id;

-- name: ListSubTypologyNames :many
SELECT DISTINCT(name) FROM sub_typology ORDER BY name ASC;
