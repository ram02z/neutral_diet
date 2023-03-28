-- name: GetAggregateFoodItem :one
SELECT
    *
FROM
    aggregate_food_item
WHERE
    food_item_id = $1;

-- name: GetRegionalAggregateFoodItem :one
SELECT
    *
FROM
    regional_aggregate_food_item
WHERE
    food_item_id = $1
    AND region = $2;

-- name: ListAggregateFoodItems :many
SELECT
    a.food_item_id AS id,
    f.name AS food_name,
    t.name AS typology_name,
    s.name AS sub_typology_name,
    a.n,
    ROUND(a.median_carbon_footprint, 3)::decimal AS median_carbon_footprint
FROM
    aggregate_food_item a
    INNER JOIN food_item f ON a.food_item_id = f.id
    INNER JOIN typology t ON f.typology_id = t.id
    LEFT JOIN sub_typology s ON t.sub_typology_id = s.id;

-- name: ListAggregateFoodItemsByRegion :many
SELECT
    a.food_item_id AS id,
    f.name AS food_name,
    t.name AS typology_name,
    s.name AS sub_typology_name,
    a.n,
    ROUND(a.median_carbon_footprint, 3)::decimal AS median_carbon_footprint
FROM
    regional_aggregate_food_item a
    INNER JOIN food_item f ON a.food_item_id = f.id
    INNER JOIN typology t ON f.typology_id = t.id
    LEFT JOIN sub_typology s ON t.sub_typology_id = s.id
WHERE
    a.region = $1;

-- name: ListAggregateTypologiesByRegion :many
SELECT
    t.id AS typology_id,
    COUNT(*) AS n,
    ROUND(PERCENTILE_CONT(0.50) WITHIN GROUP (ORDER BY l.carbon_footprint), 3)::decimal AS median_carbon_footprint
FROM
    life_cycle l
    INNER JOIN food_item f ON l.food_item_id = f.id
    INNER JOIN source s ON l.source_id = s.id
    INNER JOIN typology t ON f.typology_id = t.id
WHERE
    s.region = $1
GROUP BY
    t.id;

-- name: ListAggregateSubTypologiesByRegion :many
SELECT
    st.id AS sub_typology_id,
    COUNT(*) AS n,
    ROUND(PERCENTILE_CONT(0.50) WITHIN GROUP (ORDER BY l.carbon_footprint), 3)::decimal AS median_carbon_footprint
FROM
    life_cycle l
    INNER JOIN food_item f ON l.food_item_id = f.id
    INNER JOIN source s ON l.source_id = s.id
    INNER JOIN typology t ON f.typology_id = t.id
    INNER JOIN sub_typology st ON t.sub_typology_id = st.id
WHERE
    s.region = $1
GROUP BY
    st.id;
