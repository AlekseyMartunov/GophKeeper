CREATE TABLE IF NOT EXISTS cards (
    card_id serial PRIMARY KEY,
    card_name VARCHAR (200) NOT NULL,
    card_number VARCHAR (16) NOT NULL,
    owner VARCHAR (20) NOT NULL,
    cvv VARCHAR (3) NOT NULL,
    card_date VARCHAR (4) NOT NULL,
    created_time TIMESTAMP NOT NULL,
    fk_user_id INTEGER REFERENCES users(user_id) NOT NULL,
    UNIQUE (pair_name, fk_user_id)
    );