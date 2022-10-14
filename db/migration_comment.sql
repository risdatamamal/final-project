-- creating comment table
create table comments(
    id SERIAL NOT NULL PRIMARY KEY,
    user_id int,
    photo_id int,
    message text,
    created_at date,
    updated_at date,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (photo_id) REFERENCES photos(id)
);
--

-- JOIN USERS ID TO COMMENTS
select *
from comments o 
	join users u on o.user_id = u.id;
--

-- JOIN PHOTOS ID TO COMMENTS
select *
from comments o 
	join photos u on o.photo_id = u.id;
--