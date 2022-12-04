CREATE TABLE IF NOT EXISTS "food_item_log" (
    "id" serial PRIMARY KEY,
    "food_item_id" int NOT NULL REFERENCES "food_item" ("id") ON DELETE CASCADE,
    "weight" decimal NOT NULL,
    "carbon_footprint" decimal NOT NULL
);
