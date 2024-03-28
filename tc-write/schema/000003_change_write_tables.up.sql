DROP TABLE year;

DROP TABLE month;

DROP TABLE day;

DROP TABLE activity;

CREATE TABLE activity_event_type
(
    id serial not null unique,
    event_type varchar(255) not null
);

CREATE TABLE activity_event
(
    id uuid DEFAULT uuid_generate_v4 (),
    employee_id int  not null,
    check_duration bigint not null,
    check_time timestamp not null,
    was_active boolean not null,
    event_type_id int references activity_event_type (id) on delete cascade not null
);
