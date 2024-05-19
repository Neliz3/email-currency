CREATE TABLE users
(
    id serial not null PRIMARY KEY,
    email varchar(255) not null unique
);