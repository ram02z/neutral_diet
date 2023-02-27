CREATE MATERIALIZED VIEW IF NOT EXISTS regional_aggregate_food_item AS
SELECT
    f.id AS food_item_id,
    s.region_name,
    COUNT(*) AS n,
    CAST(PERCENTILE_CONT(0.50) WITHIN GROUP (ORDER BY l.carbon_footprint) AS DECIMAL) AS median_carbon_footprint
FROM
    life_cycle l
    INNER JOIN food_item f ON l.food_item_id = f.id
    INNER JOIN source s ON l.source_id = s.id
GROUP BY
    f.id,
    s.region_name WITH DATA;

CREATE UNIQUE INDEX IF NOT EXISTS regional_aggregate_food_item_idx
  ON regional_aggregate_food_item (food_item_id, region_name);
