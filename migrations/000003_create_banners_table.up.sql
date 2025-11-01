CREATE TABLE IF NOT EXISTS banners(
   id serial PRIMARY KEY,
   banner_name VARCHAR(255),
   banner_image TEXT,
   description TEXT,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP 
);