BEGIN;
CREATE OR REPLACE FUNCTION update_updated_at_time ()
    RETURNS TRIGGER
    AS $$
BEGIN
    NEW.updated_at = COALESCE(NEW.updated_at, now());
    RETURN NEW;
END;
$$
LANGUAGE 'plpgsql';
COMMIT;
