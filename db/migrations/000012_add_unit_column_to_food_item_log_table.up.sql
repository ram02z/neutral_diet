DROP TYPE IF EXISTS weight_unit;

CREATE TYPE weight_unit AS ENUM (
    'kilogram',
    'gram',
    'ounce',
    'pound'
);

ALTER TABLE "food_item_log"
  ADD COLUMN weight_unit weight_unit NOT NULL DEFAULT 'kilogram';
