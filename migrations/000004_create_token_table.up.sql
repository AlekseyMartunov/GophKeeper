CREATE TABLE IF NOT EXISTS tokens (
    token_id serial PRIMARY KEY,
    token_name VARCHAR (200) NOT NULL,
    token VARCHAR (200) NOT NULL,
    created_time TIMESTAMP NOT NULL,
    is_blocked BOOLEAN DEFAULT FALSE,
    external_user_id VARCHAR (200) NOT NULL,
    fk_user_id INTEGER REFERENCES users(user_id) NOT NULL,
    UNIQUE (token_name, fk_user_id)
    );