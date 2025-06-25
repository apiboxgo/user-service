-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id   uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    email VARCHAR(120) NOT NULL UNIQUE,
    password VARCHAR(60) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP  NULL DEFAULT NULL,
    deleted_at TIMESTAMP  NULL DEFAULT NULL

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users
-- +goose StatementEnd
