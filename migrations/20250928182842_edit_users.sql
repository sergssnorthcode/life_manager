-- +goose Up
ALTER TABLE users
RENAME COLUMN second_name TO last_name;

-- +goose Down
ALTER TABLE users
RENAME COLUMN last_name TO second_name;