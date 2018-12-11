create database annotations;

CREATE EXTENSION citext;

create index author_index on annotations(author_id);
create index search_index on annotations(subject, created);

create table anntotations (
    id serial,
    author_id integer,
    subject citext,
    files text[],
    created TIMESTAMP,

)