CREATE TABLE users (
    id SERIAL NOT NULL UNIQUE,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    registered_at DATE NULL,
    last_visit_at DATE NULL
);
CREATE TABLE wallets (
    id SERIAL NOT NULL UNIQUE,
    balance FLOAT NOT NULL
);
CREATE TABLE users_and_wallets (
    id SERIAL NOT NULL UNIQUE,
    user_id INT REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    wallet_id INT REFERENCES wallets (id) ON DELETE CASCADE NOT NULL
);