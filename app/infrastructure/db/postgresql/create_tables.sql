create table if not exists public.users
(
    id              uuid     not null
        constraint users_pk
            primary key,
    username        char(32) not null,
    email           char(50) not null,
    hashed_password char(70) not null
);

create table if not exists public.balance
(
    id      uuid              not null
        constraint balance_pk
            primary key,
    user_id uuid              not null
        constraint balance_users_id_fk
            references public.users,
    amount  integer default 0 not null
        constraint check_amount
            check (amount >= 0)
);
