create table companies
(
    id serial primary key,
    uuid             varchar,
    name             varchar,
    description      text,
    employees_amount integer,
    registered       boolean,
    type             varchar
);

alter table companies
    owner to postgres;

