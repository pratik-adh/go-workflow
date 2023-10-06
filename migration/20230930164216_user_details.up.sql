CREATE TABLE user_details (
   id int primary key,
   name  varchar(50) NOT NULl UNIQUE,
   age int,
   email  varchar(50) UNIQUE,
   phone varchar(50) UNIQUE,
   address  varchar(50)
);