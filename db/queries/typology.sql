-- name: CreateTypology :one
INSERT INTO typology (name, sub_typologies)
    VALUES ($1, $2)
RETURNING
    id;
