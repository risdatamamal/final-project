-- creating photo table
create table photos(
    id SERIAL NOT NULL PRIMARY KEY,
    title text,
    caption text,
    photo_url text,
    user_id int,
    created_at date,
    updated_at date,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
--

-- JOIN USERS ID TO PHOTOS
select *
from photos o 
	join users u on o.user_id = u.id;
--