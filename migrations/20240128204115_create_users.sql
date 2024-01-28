-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
	id bigserial not null primary key,
	email varchar not null unique,
	encrypted_password varchar not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
DROP TABLE users;
-- +goose StatementEnd

-- goose -dir migrations postgres "user=app password=passwd dbname=restapi_dev sslmode=disable" up 20240128204115_create_users.sql