create table if not exists users
(
    id serial primary key,
    name text not null
);

create table if not exists books
(
    id serial primary key,
    name text,
    user_id int references users(id) not null
--     foreign key(user_id) references users(id)
);
