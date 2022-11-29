BEGIN;
CREATE TABLE IF NOT EXISTS "food_item" (
    "id" serial PRIMARY KEY,
    "name" text NOT NULL,
    "carbon_footprint" decimal NOT NULL,
    "typology_id" int NOT NULL REFERENCES "typology" ("id") ON DELETE CASCADE,
    "source_id" int NOT NULL REFERENCES "source" ("id") ON DELETE CASCADE
);
COMMIT;
