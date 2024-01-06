-- +goose Up
ALTER TABLE note ADD COLUMN tag text;

-- +goose Down
ALTER TABLE note DROP COLUMN tag;
