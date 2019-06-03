
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE app_user (id SERIAL PRIMARY KEY, username VARCHAR(255));

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE app_user;
