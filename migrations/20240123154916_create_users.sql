-- +goose Up
-- +goose StatementBegin
create table users
(
    id       serial not null primary key,
    uuid text unique,
    email    text not null unique,
    pass_hash bytea not null,
    role text not null,
    created_at timestamp not null,
    updated_at timestamp,
    logged_in timestamp,
    logged_out timestamp
);

alter table users owner to postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd