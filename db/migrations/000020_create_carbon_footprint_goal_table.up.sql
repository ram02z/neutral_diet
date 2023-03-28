CREATE TABLE IF NOT EXISTS "carbon_footprint_goal" (
  "id" serial PRIMARY KEY,
  "user_id" int NOT NULL REFERENCES "user" ("id") ON DELETE CASCADE,
  "description" text NOT NULL,
  "start_date" DATE NOT NULL,
  "end_date" DATE NOT NULL,
  "start_carbon_footprint" decimal NOT NULL,
  "target_carbon_footprint" decimal NOT NULL,
  "completed" boolean NOT NULL DEFAULT FALSE
)
