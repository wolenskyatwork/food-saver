
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE activity_name (id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE activity_name;
