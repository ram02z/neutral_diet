ALTER TABLE "food_item_log"
    ADD COLUMN "created_at" timestamptz NOT NULL DEFAULT NOW();

ALTER TABLE "food_item_log"
    ADD COLUMN "updated_at" timestamptz NOT NULL DEFAULT NOW();

CREATE TRIGGER update_updated_at_time_trigger
    BEFORE UPDATE ON "food_item_log"
    FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at_time ();
