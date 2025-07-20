-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS record (
    id INTEGER PRIMARY KEY,
    title varchar(255),
    description varchar(255),
    release_date date,
    duration decimal(10,2)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS record;
-- +goose StatementEnd