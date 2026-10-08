CREATE TABLE users (
    id bigserial not null primary key,
    username varchar not null,
    email varchar not null,
    hash_password varchar not null,
    coins bigint default 1000
);

CREATE TABLE merch (
    id bigserial not null primary key,
    name varchar not null,
    price bigint not null
);

INSERT INTO merch (name, price) VALUES
    ('t-shirt', 80),
    ('cup', 20),
    ('book', 50),
    ('pen', 10),
    ('powerbank', 200),
    ('hoody', 300),
    ('umbrella', 200),
    ('socks', 10),
    ('wallet', 50),
    ('pink-hoody', 500);

CREATE TABLE purchases (
    id bigserial not null primary key,
    user_id bigint references users(id),
    merch_id bigint references merch(id),
    created_at timestamptz not null default now()
);

CREATE TABLE coins_transfers (
    id bigserial not null primary key,
    sender_id bigint references users(id),
    receiver_id bigint references users(id),
    amount bigint not null
);