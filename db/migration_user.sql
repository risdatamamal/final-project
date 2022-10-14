-- creating user table 
CREATE TABLE users(
    id SERIAL NOT NULL PRIMARY KEY,
    username text,
    email text,
    password text,
    dbo date,
    age int,
    created_at date,
    updated_at date
);