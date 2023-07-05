create table if not exists offices
(
    uuid       uuid not null
    constraint office_pk
    primary key,
    name       text,
    address    text,
    created_at timestamp
);
create table if not exists users
(
    uuid        uuid not null
    constraint user_pk
    primary key,
    name        text,
    office_uuid uuid
    constraint user_office_uuid_fk
    references public.offices
    on update cascade on delete cascade,
    created_at  timestamp
);
create table if not exists orders
(
    uuid       uuid not null
    constraint order_pk
    primary key,
    user_uuid  uuid
    constraint order_users_uuid_fk
    references public.users
    on update cascade on delete cascade,
    created_at timestamp
);
create table if not exists order_item
(
    order_uuid   uuid
    constraint order_item_order_uuid_fk
    references public.orders
    on update cascade on delete cascade,
    count        integer,
    product_uuid uuid
);


