CREATE TABLE registration_details (
   id varchar(50) PRIMARY KEY,
   name  varchar(50) NOT NULl Unique,
   email  varchar(50) NOT NULl Unique,
   password  varchar(50) NOT NULl,
   phone varchar(20) NOT NULl,
   age int NOT NULL,
   address  varchar(50) NOT NULl
);
