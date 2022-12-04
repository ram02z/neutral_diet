CREATE MATERIALIZED VIEW IF NOT EXISTS aggregate_food_item AS
SELECT
    f.id AS food_item_id,
    COUNT(*) AS n,
    CAST(PERCENTILE_CONT(0.50) WITHIN GROUP (ORDER BY l.carbon_footprint) AS DECIMAL) AS median_carbon_footprint
FROM
    life_cycle l
    INNER JOIN food_item f ON l.food_item_id = f.id
GROUP BY
    f.id WITH DATA;
