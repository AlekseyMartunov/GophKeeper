CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    user_id serial PRIMARY KEY,
    login VARCHAR (100) UNIQUE NOT NULL,
    password VARCHAR (100) NOT NULL,
    external_id  uuid DEFAULT uuid_generate_v4()
    );