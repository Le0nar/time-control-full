CREATE TABLE activities
(
    id UUID not null,
    was_active boolean not null,
    check_duration bigint not null,
    employee_id int references employees (id) on delete cascade not null,
    check_time timestamp not null
);
