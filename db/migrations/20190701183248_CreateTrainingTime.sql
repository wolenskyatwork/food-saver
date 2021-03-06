
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE training_time (id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL, meals INTEGER NOT NULL);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE training_time;
