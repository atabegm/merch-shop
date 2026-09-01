-- Инициализация схемы БД (заглушка под docker-compose на Mac).
CREATE TABLE users(
    id bigserial not null primary key,
    username varchar not null,
    email varchar not null,
    hash_password varchar not null,    
    coins bigint default 1000
);

INSERT INTO users(username, email, hash_password, coins) VALUES (

)

CREATE TABLE merch (
    id bigserial not null primary key,
    name varchar not null,
    price bigint not null
); 

INSERT INTO merch (name, price) VALUES 
    ('t_shirt', 80), 
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
    price bigint not null  
);

CREATE TABLE coins_transfers (
    id bigserial not null primary key,
    sender_id bigint references users(id),
    receiver_id bigint references users(id),
    amount bigint not null
);

CREATE TABLE purchases_merch (
    purchase_id bigint references purchases(id),
    merch_id bigint references merch(id),
    amount bigint not null
);

-- users:
-- (1, ibrahim, 100)
-- (2, muhammad, 1900)

-- coins_transfer:
-- (1, 1, 2, 300) // 700, 1300
-- (1, 2, 1, 400) // 1100, 900
-- (1, 1, 2, 1000) // 100, 1900