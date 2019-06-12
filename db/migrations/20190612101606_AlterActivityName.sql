
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE activity_name ADD app_user_id int REFERENCES app_user(id) NOT NULL DEFAULT(1);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE activity_name DROP COLUMN app_user_id;
