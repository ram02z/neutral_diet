DROP TYPE IF EXISTS weight_unit;

ALTER TABLE "food_item_log"
    DROP COLUMN IF EXISTS "weight_unit";
