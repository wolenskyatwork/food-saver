
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE training_day (
  id SERIAL PRIMARY KEY,
  training_level_id INTEGER REFERENCES training_level(id),
  cut_id INTEGER REFERENCES cut_level(id),
  training_time_id INTEGER REFERENCES training_time(id),
  app_user_id INTEGER REFERENCES app_user(id),
  datetime_completed DATE
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE training_day
