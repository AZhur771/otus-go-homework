-- +goose Up
ALTER TABLE events
    ADD COLUMN sent BOOLEAN NOT NULL DEFAULT FALSE;

-- +goose Down
ALTER TABLE events
    DROP COLUMN sent;