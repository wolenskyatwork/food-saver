
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied


CREATE TABLE weight (id SERIAL PRIMARY KEY, app_user_id INTEGER REFERENCES app_user(id), weight DECIMAL, weight_date DATE);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE weight;
