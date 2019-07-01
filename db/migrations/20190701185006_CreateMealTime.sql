
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE meal_time (
  id SERIAL PRIMARY KEY,
  meal_number INTEGER,
  time_completed TIME,
  training_day_id INTEGER REFERENCES training_day(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE meal_time;
