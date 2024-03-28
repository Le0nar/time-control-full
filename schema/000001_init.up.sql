CREATE TABLE day_activity
(
    id serial not null unique,
    employee_id int  not null,
    activity_date date  not null,
    activity_time bigint not null
);
