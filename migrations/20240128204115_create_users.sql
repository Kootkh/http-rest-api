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

-- goose -dir migrations postgres "host=192.168.30.10 user=app_restapi_dev password=passwd dbname=restapi_dev sslmode=disable" up 20240128204115_create_users.sql
-- goose -dir migrations postgres "host=192.168.30.10 user=app_restapi_dev password=passwd dbname=restapi_dev sslmode=disable" down 20240128204115_create_users.sql