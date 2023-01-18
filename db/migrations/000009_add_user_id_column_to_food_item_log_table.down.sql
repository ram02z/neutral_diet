ALTER TABLE "food_item_log"
    DROP CONSTRAINT IF EXISTS "food_item_log_user_id_fkey" CASCADE;

ALTER TABLE "food_item_log"
    DROP COLUMN IF EXISTS "user_id";
