-- +goose Up
ALTER TABLE post ADD COLUMN visible BOOLEAN NOT NULL DEFAULT FALSE;

-- +goose Down
ALTER TABLE post DROP COLUMN visible;
