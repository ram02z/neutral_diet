CREATE TABLE IF NOT EXISTS "region" (
    "name" text PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS "source" (
    "id" serial PRIMARY KEY,
    "reference" text UNIQUE NOT NULL,
    "year" smallint NOT NULL,
    "region" text NOT NULL REFERENCES "region" ("name") ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "sub_typology" (
    "id" serial PRIMARY KEY,
    "name" text UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS "typology" (
    "id" serial PRIMARY KEY,
    "name" text UNIQUE NOT NULL,
    "description" text NOT NULL,
    "sub_typology_id" int REFERENCES "sub_typology" ("id") ON DELETE CASCADE
);

DROP TYPE IF EXISTS cf_types;

CREATE TYPE cf_types AS ENUM (
    'typology',
    'sub_typology',
    'item'
);

CREATE TABLE IF NOT EXISTS "food_item" (
    "id" serial PRIMARY KEY,
    "name" text NOT NULL,
    "typology_id" int NOT NULL REFERENCES "typology" ("id") ON DELETE CASCADE,
    "suggested_cf" cf_types NOT NULL
);

CREATE TABLE IF NOT EXISTS "life_cycle" (
    "id" serial PRIMARY KEY,
    "carbon_footprint" decimal NOT NULL,
    "food_item_id" int NOT NULL REFERENCES "food_item" ("id") ON DELETE CASCADE,
    "source_id" int NOT NULL REFERENCES "source" ("id") ON DELETE CASCADE
);
