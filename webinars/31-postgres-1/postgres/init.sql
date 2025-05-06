create table if not exists users
(
    id serial primary key,
    name text not null
);

create table if not exists authors
(
    id serial primary key,
    name text,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);

create table if not exists books
(
    id serial primary key,
    name text,
    user_id int references users(id),
    author_id int references authors(id) not null
);
