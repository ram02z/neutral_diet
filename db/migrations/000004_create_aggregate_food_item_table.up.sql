BEGIN;
CREATE TABLE IF NOT EXISTS aggregate_food_item AS
SELECT
    name,
    COUNT(*) AS n,
    CAST(PERCENTILE_CONT(0.50) WITHIN GROUP (ORDER BY carbon_footprint) AS DECIMAL) AS median_carbon_footprint,
    typology_id
FROM
    food_item
GROUP BY
    name,
    typology_id;
ALTER TABLE aggregate_food_item
    ADD COLUMN id serial PRIMARY KEY;
COMMIT;
