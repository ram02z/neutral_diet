BEGIN;
CREATE TABLE IF NOT EXISTS "typology" (
    "id" serial PRIMARY KEY,
    "name" text UNIQUE NOT NULL,
    "sub_typologies" text[]
);
COMMIT;
