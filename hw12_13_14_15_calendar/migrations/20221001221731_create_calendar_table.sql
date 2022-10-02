-- +goose Up
CREATE TABLE users (
   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
   username VARCHAR(100) NOT NULL
);

CREATE TABLE events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    date_start TIMESTAMP NOT NULL,
    duration INTERVAL NOT NULL,
    description VARCHAR(255) NOT NULL,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    notification_period INTERVAL NOT NULL
);

create index user_id_idx on events (user_id);
create index date_start_idx on events using btree (date_start);

INSERT INTO users (username)
VALUES ('anonymous');

-- +goose Down
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS users;
