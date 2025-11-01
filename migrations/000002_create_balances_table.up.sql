CREATE TABLE IF NOT EXISTS balances(
   id serial PRIMARY KEY,
   user_id INT NOT NULL,
   amount DOUBLE PRECISION NOT NULL DEFAULT 0,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP 
);