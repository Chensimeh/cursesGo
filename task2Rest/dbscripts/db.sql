CREATE TABLE users (
    id INT UNIQUE NOT NULL,
    firstname  varchar(255),
    lastname  varchar(255),
    email varchar(255)
);
INSERT INTO users(id,firstname,lastname,email) values(1,'Tolia','Picus','tarakan@net.net');
SELECT * FROM users;
