BEGIN;
CREATE OR REPLACE FUNCTION update_aggregate_food_item ()
    RETURNS TRIGGER
    AS $$
BEGIN
    PERFORM
        1
    FROM
        aggregate_food_item
    WHERE
        aggregate_food_item.name = NEW.name
        AND aggregate_food_item.typology_id = NEW.typology_id
    LIMIT 1;
    IF FOUND THEN
        UPDATE
            aggregate_food_item
        SET
            n = subquery.n,
            median_carbon_footprint = subquery.median_carbon_footprint
        FROM (
            SELECT
                COUNT(*) AS n,
                CAST(PERCENTILE_CONT(0.50) WITHIN GROUP (ORDER BY carbon_footprint) AS DECIMAL) AS median_carbon_footprint
            FROM
                food_item
            WHERE
                food_item.name = NEW.name
                AND food_item.typology_id = NEW.typology_id GROUP BY
                    food_item.name,
                    food_item.typology_id) AS subquery
    WHERE
        aggregate_food_item.name = NEW.name
            AND aggregate_food_item.typology_id = NEW.typology_id;
    ELSE
        INSERT INTO aggregate_food_item (name, n, median_carbon_footprint, typology_id)
            VALUES (NEW.name, 1, NEW.carbon_footprint, NEW.typology_id);
    END IF;
    RETURN NEW;
END;
$$
LANGUAGE plpgsql;
CREATE TRIGGER food_item_audit AFTER INSERT ON food_item FOR EACH ROW EXECUTE FUNCTION update_aggregate_food_item ();
COMMIT;
