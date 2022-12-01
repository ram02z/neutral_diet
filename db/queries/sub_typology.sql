-- name: CreateSubTypology :one
INSERT INTO sub_typology (name)
    VALUES ($1)
RETURNING
    id;
