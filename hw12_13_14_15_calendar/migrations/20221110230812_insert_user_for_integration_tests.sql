-- +goose Up
INSERT INTO users (id, username)
VALUES ('19a12b49-a57a-4f1e-8e66-152be08e6165', 'user_for_tests');

-- +goose Down
DELETE FROM users WHERE id = '19a12b49-a57a-4f1e-8e66-152be08e6165';
