-- name: CreateSource :one
INSERT INTO source (reference, year, region)
    VALUES ($1, $2, $3)
RETURNING
    id;

-- name: ListSourcesByFoodItem :many
SELECT
    s.reference,
    s.year,
    s.region
FROM
    life_cycle l
    INNER JOIN source s ON l.source_id = s.id
WHERE
    l.food_item_id = $1
GROUP BY
    s.id;
