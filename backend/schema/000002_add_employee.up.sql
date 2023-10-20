CREATE TABLE employees
(
    id serial not null unique,
    email varchar(255) not null unique,
    password_hash varchar(255) not null,
    first_name varchar(255) not null,
    second_name varchar(255) not null,
    patronymic varchar(255),
    company_id int references companies (id) on delete cascade not null
);
