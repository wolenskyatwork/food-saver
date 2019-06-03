
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE activity (id SERIAL PRIMARY KEY, activity_id INTEGER REFERENCES activity_name(id), app_user_id INTEGER REFERENCES app_user(id), datetime_completed DATE);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE activity;

