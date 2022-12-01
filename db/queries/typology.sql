-- name: CreateTypology :one
INSERT INTO typology (name, description, sub_typology_id)
    VALUES ($1, $2, $3)
RETURNING
    id;
