BEGIN;
DROP TRIGGER IF EXISTS food_item_audit ON food_item;
DROP FUNCTION IF EXISTS update_aggregate_food_item;
COMMIT;
