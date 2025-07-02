-- +goose Up
-- +goose StatementBegin
create table if not exists users
(
    id serial primary key,
    name text not null
);
create table if not exists device_types
(
    id serial primary key,
    name text not null
);
create table if not exists devices
(
    id serial primary key,
    name text not null,
    type_id int references device_types(id) not null,
    user_id int references users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users cascade;
drop table device_types cascade;
drop table devices;
-- +goose StatementEnd
