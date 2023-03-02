-- name: CreateFoodItem :one
INSERT INTO food_item (name, typology_id, suggested_cf)
    VALUES ($1, $2, $3)
RETURNING
    id;

-- name: GetFoodItemInfo :one
SELECT
    t.name AS typology_name,
    st.name AS sub_typology_name,
    COUNT(l.food_item_id) AS n
FROM
    life_cycle l
    INNER JOIN food_item f ON l.food_item_id = f.id
    INNER JOIN typology t ON f.typology_id = t.id
    LEFT JOIN sub_typology st ON t.sub_typology_id = st.id
WHERE
    l.food_item_id = $1
GROUP BY
    t.name,
    st.name;

-- name: GetFoodItemInfoByRegion :one
SELECT
    t.name AS typology_name,
    st.name AS sub_typology_name,
    COUNT(l.food_item_id) AS n
FROM
    life_cycle l
    INNER JOIN source s ON l.source_id = s.id
    INNER JOIN food_item f ON l.food_item_id = f.id
    INNER JOIN typology t ON f.typology_id = t.id
    LEFT JOIN sub_typology st ON t.sub_typology_id = st.id
WHERE
    l.food_item_id = $1
    AND s.region = $2
GROUP BY
    t.name,
    st.name;
