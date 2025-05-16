select 	'["a", {"b":1}]'::jsonb #- '{0}';
select 	'["a", {"b":1}]'::jsonb #- '{1}';
select 	'["a", {"b":1}]'::jsonb #- '{1,b}';

insert into device_types(name) values ('phone'), ('laptop'), ('watches');

insert into users(name) values ('John'), ('Peter');

insert into devices(name, type_id, user_id, attributes)
values ('Lenovo Yoga Pro 1i', 2, 2, '{"color": "gray", "weight": "500g", "HDD": "1TB SSD", "RAM": "32GB"}');
insert into devices(name, type_id, user_id, attributes)
values ('Lenovo Yoga Pro 2i', 2, 2, '{"color": "gray", "weight": "500g", "HDD": "1TB SSD", "RAM": "32GB"}');
insert into devices(name, type_id, user_id, attributes)
values ('Lenovo Yoga Pro 3i', 2, 2, '{"color": "gray", "weight": "500g", "HDD": "1TB SSD", "RAM": "32GB"}');
insert into devices(name, type_id, user_id, attributes)
values ('Lenovo Yoga Pro 4i', 2, 2, '{"color": "gray", "weight": "500g", "HDD": "1TB SSD", "RAM": "32GB"}');
insert into devices(name, type_id, user_id, attributes)
values ('Lenovo Yoga Pro 5i', 2, 2, '{"color": "gray", "weight": "500g", "HDD": "1TB SSD", "RAM": "32GB"}');
insert into devices(name, type_id, user_id, attributes)
values ('Lenovo Yoga Pro 6i', 2, 2, '{"color": "gray", "weight": "500g", "HDD": "1TB SSD", "RAM": "32GB"}');
insert into devices(name, type_id, user_id, attributes)
values ('Lenovo Yoga Pro 8i', 2, 2, '{"color": "gray", "weight": "500g", "HDD": "1TB SSD", "RAM": "32GB"}');
insert into devices(name, type_id, user_id, attributes)
values ('Lenovo Yoga Pro 9i', 2, 2, '{"color": "gray", "weight": "500g", "HDD": "1TB SSD", "RAM": "32GB"}');
insert into devices(name, type_id, user_id, attributes)
values ('Lenovo Yoga Pro 10i', 2, 2, '{"color": "gray", "weight": "500g", "HDD": "1TB SSD", "RAM": "32GB"}');

INSERT INTO devices(name, type_id, user_id, attributes)
SELECT data->>'name', (data->>'type_id')::integer, (data->>'user_id')::integer, data->'attributes'
FROM temp;

-- explain analyse
select devices.name,
       devices.attributes
from devices
where attributes @> '{"HDD": "1TB SSD"}';

create index device_attributes_idx on devices using GIN (attributes);
drop index device_attributes_idx;

CREATE TABLE temp (data jsonb);
COPY temp (data) FROM './laptops.json';

create index test_data_idx on temp using GIN (data);

drop index test_data_idx;

select users.name, devices.name, count(devices.name)
from users
         join devices on devices.user_id = users.id
where devices.attributes @> '{"HDD": "1TB SSD"}'
group by users.name, devices.name
having count(devices.name) > 15;
