-- +goose Up
-- +goose StatementBegin
insert into users (uuid, email, pass_hash, role, created_at)
values ('cae96412-2bb2-4900-a3cd-96f465ba4db2', 'example@mail.com', 'hash', 'EMPLOYEE', '2024-01-24T03:58:03');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
truncate table users;
-- +goose StatementEnd
