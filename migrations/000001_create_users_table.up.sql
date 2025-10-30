CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   first_name VARCHAR (50),
   last_name VARCHAR (50),
   password VARCHAR (50) NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP 
);