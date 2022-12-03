ALTER TABLE typology
    DROP CONSTRAINT IF EXISTS unique_typology;

ALTER TABLE typology
    ADD CONSTRAINT unique_typology UNIQUE (name, sub_typology_id);
