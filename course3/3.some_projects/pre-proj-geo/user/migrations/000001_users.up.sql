CREATE TABLE IF NOT EXISTS users (
    username VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    user_role VARCHAR(255)
);