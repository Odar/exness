-- auto-generated definition
create type transaction_type as enum ('replenish', 'transfer');

alter type transaction_type owner to postgres;

-- auto-generated definition
create sequence account_id_seq;

alter sequence account_id_seq owner to postgres;

-- auto-generated definition
create sequence transaction_id_seq;

alter sequence transaction_id_seq owner to postgres;

-- auto-generated definition
create table account
(
    id         bigserial                          not null
        constraint account_pk
            primary key,
    cents      bigint                   default 0 not null,
    created_at timestamp with time zone default timezone('utc'::text, now())
);

alter table account
    owner to postgres;

create unique index account_id_uindex
    on account (id);

-- auto-generated definition
create table transaction
(
    id              bigserial                                                     not null
        constraint transaction_pk
            primary key,
    from_account_id bigint
        constraint transaction_account_from_id_fk
            references account
            on delete restrict,
    to_account_id   bigint                                                        not null
        constraint transaction_account_to_id_fk
            references account
            on delete restrict,
    cents           bigint                                                        not null,
    type            transaction_type                                              not null,
    committed_at    timestamp with time zone default timezone('utc'::text, now()) not null
);

alter table transaction
    owner to postgres;

create unique index transaction_id_uindex
    on transaction (id);

