-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sensors (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    address VARCHAR(255) NOT NULL,
    port INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sensors;
-- +goose StatementEnd
