-- name: ListAggregateFoodItems :many
SELECT
    a.food_item_id AS id,
    f.name AS food_name,
    t.name AS typology_name,
    s.name AS sub_typology_name,
    a.n,
    ROUND(a.median_carbon_footprint::decimal, 3) AS median_carbon_footprint
FROM
    aggregate_food_item a
    INNER JOIN food_item f ON a.food_item_id = f.id
    INNER JOIN typology t ON f.typology_id = t.id
    LEFT JOIN sub_typology s ON t.sub_typology_id = s.id;

-- name: ListAggregateFoodItemsByRegion :many
SELECT
    f.id AS food_item_id,
    COUNT(*) AS n,
    ROUND(CAST(PERCENTILE_CONT(0.50) WITHIN GROUP (ORDER BY l.carbon_footprint) AS DECIMAL), 3) AS median_carbon_footprint
FROM
    life_cycle l
    INNER JOIN food_item f ON l.food_item_id = f.id
    INNER JOIN source s ON l.source_id = s.id
WHERE
    s.region_name = $1
GROUP BY
    f.id;

-- name: ListAggregateTypologiesByRegion :many
SELECT
    t.id AS typology_id,
    COUNT(*) AS n,
    ROUND(CAST(PERCENTILE_CONT(0.50) WITHIN GROUP (ORDER BY l.carbon_footprint) AS DECIMAL), 3) AS median_carbon_footprint
FROM
    life_cycle l
    INNER JOIN food_item f ON l.food_item_id = f.id
    INNER JOIN source s ON l.source_id = s.id
    INNER JOIN typology t ON f.typology_id = t.id
WHERE
    s.region_name = $1
GROUP BY
    t.id;

-- name: ListAggregateSubTypologiesByRegion :many
SELECT
    st.id AS sub_typology_id,
    COUNT(*) AS n,
    ROUND(CAST(PERCENTILE_CONT(0.50) WITHIN GROUP (ORDER BY l.carbon_footprint) AS DECIMAL), 3) AS median_carbon_footprint
FROM
    life_cycle l
    INNER JOIN food_item f ON l.food_item_id = f.id
    INNER JOIN source s ON l.source_id = s.id
    INNER JOIN typology t ON f.typology_id = t.id
    INNER JOIN sub_typology st ON t.sub_typology_id = st.id
WHERE
    s.region_name = $1
GROUP BY
    st.id;
