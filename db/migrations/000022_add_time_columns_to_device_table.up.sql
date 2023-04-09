ALTER TABLE "device"
    ADD COLUMN "created_at" timestamptz NOT NULL DEFAULT NOW();

ALTER TABLE "device"
    ADD COLUMN "updated_at" timestamptz NOT NULL DEFAULT NOW();

CREATE TRIGGER update_updated_at_time_trigger
    BEFORE UPDATE ON "device"
    FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at_time ();
