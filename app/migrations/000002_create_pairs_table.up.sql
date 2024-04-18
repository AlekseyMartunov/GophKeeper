CREATE TABLE IF NOT EXISTS pairs (
    pair_id serial PRIMARY KEY,
    pair_name VARCHAR (200) NOT NULL,
    password VARCHAR (200) NOT NULL,
    login VARCHAR (200) NOT NULL,
    created_time TIMESTAMP NOT NULL,
    fk_user_id INTEGER REFERENCES users(user_id) NOT NULL,
    UNIQUE (pair_name, fk_user_id)
    );