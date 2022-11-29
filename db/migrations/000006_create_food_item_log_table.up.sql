BEGIN;
CREATE TABLE IF NOT EXISTS "food_item_log" (
    "id" serial PRIMARY KEY,
    "aggregate_food_item_id" int NOT NULL REFERENCES "aggregate_food_item" ("id") ON DELETE CASCADE,
    "weight" decimal NOT NULL,
    "carbon_footprint" decimal NOT NULL
);
COMMIT;
