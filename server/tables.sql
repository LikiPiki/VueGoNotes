-- if change schema rebild docker compose

CREATE TABLE IF NOT EXISTS users (
    id SERIAL,
    username TEXT UNIQUE NOT NULL,
    password text NOT NULL,
    is_admin BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS notes (
    id SERIAL,
    title TEXT NOT NULL,
    content TEXT NOT NULL
);