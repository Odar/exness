create database exness;

create table account
(
    id bigserial not null,
    balance bigint default 0 not null
    created_at timestamp default now() not null
);

create unique index account_id_uindex
	on account (id);

alter table account
    add constraint account_pk
        primary key (id);

create table transaction
(
    id bigserial not null,
    sender_account_id bigint not null,
    recipient_account_id bigint not null,
    cents bigint not null,
    type varchar not null,
    committed_in timestamp default now() not null
);

create unique index transaction_id_uindex
	on transaction (id);

alter table transaction
    add constraint transaction_pk
        primary key (id);


