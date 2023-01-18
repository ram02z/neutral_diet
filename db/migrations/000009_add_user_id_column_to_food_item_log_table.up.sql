ALTER TABLE "food_item_log"
    ADD COLUMN user_id INT NOT NULL;

ALTER TABLE "food_item_log"
    ADD CONSTRAINT food_item_log_user_id_fkey FOREIGN KEY (user_id) REFERENCES "user" (id) ON DELETE CASCADE;
