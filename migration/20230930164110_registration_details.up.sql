CREATE TABLE registration_details (
   id varchar(36) PRIMARY KEY,
   name  varchar(50) NOT NULl UNIQUE,
   email  varchar(50) NOT NULl UNIQUE,
   password  varchar(50) NOT NULl,
   phone varchar(20) NOT NULl,
   age int NOT NULL,
   address  varchar(50) NOT NULl
);
