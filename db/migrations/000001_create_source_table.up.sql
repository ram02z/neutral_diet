BEGIN;
CREATE TABLE IF NOT EXISTS "source" (
    "id" serial PRIMARY KEY,
    "reference" text UNIQUE NOT NULL,
    "year" smallint,
    "location" text
);
COMMIT;
