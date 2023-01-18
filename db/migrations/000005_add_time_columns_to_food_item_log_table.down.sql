ALTER TABLE "food_item_log"
    DROP COLUMN IF EXISTS "created_at";

ALTER TABLE "food_item_log"
    DROP COLUMN IF EXISTS "updated_at";

DROP TRIGGER IF EXISTS update_updated_at_time_trigger ON "food_item_log";
