ALTER TABLE "device"
    DROP COLUMN IF EXISTS "created_at";

ALTER TABLE "device"
    DROP COLUMN IF EXISTS "updated_at";

DROP TRIGGER IF EXISTS update_updated_at_time_trigger ON "device";
