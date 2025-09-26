-- +goose Up
CREATE TABLE
    IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        telegram_id BIGINT NOT NULL UNIQUE,
        username TEXT,
        first_name TEXT,
        second_name TEXT,
        created_at TIMESTAMPTZ NOT NULL DEFAULT now ()
    );

-- +goose Down
DROP TABLE IF EXISTS users;