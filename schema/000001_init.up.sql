CREATE TABLE activity
(
    id UUID not null,
    was_active boolean not null,
    check_duration bigint not null,
    employee_id int  not null,
    check_time timestamp not null
);
