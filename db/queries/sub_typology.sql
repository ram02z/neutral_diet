-- name: CreateSubTypology :one
INSERT INTO sub_typology (name)
    VALUES ($1)
RETURNING
    id;

-- name: ListSubTypologies :many
SELECT * FROM sub_typology;
