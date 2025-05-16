-- name: GetAllUsers :many
select id, name from users;

-- name: AddUser :exec
insert into users(name) values ($1);

-- name: GetUserDevicesCount :many
select users.name, devices.name as deviceName, count(devices.name)
from users
join devices on devices.user_id = users.id
where devices.attributes @> $1
group by users.name, devices.name
having count(devices.name)  > $2;
