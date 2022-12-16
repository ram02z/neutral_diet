CREATE TABLE IF NOT EXISTS "user" (
    "id" serial PRIMARY KEY,
    "auth_token" text UNIQUE NOT NULL,
    "region" text REFERENCES "region" ("name") ON DELETE SET NULL,
    "cf_limit" decimal,
    "created_at" timestamptz NOT NULL DEFAULT NOW(),
    "updated_at" timestamptz NOT NULL DEFAULT NOW()
);

CREATE TRIGGER update_updated_at_time_trigger
    BEFORE UPDATE ON "user"
    FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at_time ();
