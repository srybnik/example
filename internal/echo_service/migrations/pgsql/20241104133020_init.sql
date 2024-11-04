-- +goose Up
-- +goose StatementBegin

CREATE SCHEMA IF NOT EXISTS humah_resource;

CREATE EXTENSION IF NOT EXISTS pg_stat_statements WITH SCHEMA public;
CREATE EXTENSION IF NOT EXISTS pgstattuple WITH SCHEMA public;

create table if not exists humah_resource.employees
(
    employee_id      uuid                                   not null
        primary key,
    employee_name    text                                   not null,
    created_at       timestamp with time zone default now() not null,
    updated_at       timestamp with time zone default now() not null,
    constraint unique_employees
        unique (employee_id)
);

-- +goose StatementEnd
