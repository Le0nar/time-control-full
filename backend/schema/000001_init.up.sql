CREATE TABLE companies
(
    id serial not null unique,
    email varchar(255) not null unique,
    name varchar(255) not null,
    password_hash varchar(255) not null
);

-- TODO: add tables for employees and activities