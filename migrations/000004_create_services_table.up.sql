CREATE TABLE IF NOT EXISTS services(
   id serial PRIMARY KEY,
   service_code VARCHAR(255) UNIQUE,
   service_name VARCHAR(255),
   service_icon TEXT,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);