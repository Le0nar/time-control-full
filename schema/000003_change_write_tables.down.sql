CREATE TABLE year
(
    id serial not null unique,
    employee_id int  not null,
    year int  not null
);

CREATE TABLE month
(
    id serial not null unique,
    year_id int  not null,
    month int  not null
);

CREATE TABLE day
(
    id serial not null unique,
    month_id int  not null,
    day int  not null,
    activity_time bigint not null
);

CREATE TABLE activity
(
    id UUID not null,
    was_active boolean not null,
    check_duration bigint not null,
    employee_id int  not null,
    check_time timestamp not null
);
