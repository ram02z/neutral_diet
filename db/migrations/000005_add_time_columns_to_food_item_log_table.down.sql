ALTER TABLE "food_item_log"
    DROP COLUMN "created_at";

ALTER TABLE "food_item_log"
    DROP COLUMN "updated_at";

DROP TRIGGER IF EXISTS update_updated_at_time_trigger ON "food_item_log";
