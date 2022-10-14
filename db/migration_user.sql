-- creating user table 
CREATE TABLE users(
    id SERIAL NOT NULL PRIMARY KEY,
    username text,
    email text,
    password text,
    age int,
    created_at date,
    updated_at date
);