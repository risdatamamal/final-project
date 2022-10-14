-- creating social_medias table
create table social_medias(
    id SERIAL NOT NULL PRIMARY KEY,
    user_id int,
    name String,
    social_media_url String,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
--

-- JOIN USERS ID TO SOCIALMEDIAS
select *
from social_medias o 
	join users u on o.user_id = u.id;
--