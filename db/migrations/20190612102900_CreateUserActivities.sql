
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE user_activity (id SERIAL PRIMARY KEY, activity_id INTEGER REFERENCES activity_name(id), app_user_id INTEGER REFERENCES app_user(id));

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE user_activity;
