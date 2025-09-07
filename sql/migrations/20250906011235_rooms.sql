-- +goose Up
-- +goose StatementBegin
CREATE TABLE rooms (
    id INTEGER PRIMARY KEY,
    topic text NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE rooms;
-- +goose StatementEnd
