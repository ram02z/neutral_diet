CREATE TABLE IF NOT EXISTS "device" (
  "id" serial PRIMARY KEY,
  "user_id" int NOT NULL REFERENCES "user" ("id") ON DELETE CASCADE,
  "fcm_token" text UNIQUE NOT NULL
)
