CREATE TABLE IF NOT EXISTS "user" (
    "id" serial PRIMARY KEY,
    "firebase_uid" text UNIQUE NOT NULL,
    "region" int NOT NULL,
    "cf_limit" decimal NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT NOW(),
    "updated_at" timestamptz NOT NULL DEFAULT NOW()
);

CREATE TRIGGER update_updated_at_time_trigger
    BEFORE UPDATE ON "user"
    FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at_time ();
