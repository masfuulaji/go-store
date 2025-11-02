CREATE TABLE IF NOT EXISTS transactions(
   id serial PRIMARY KEY,
   invoice_number VARCHAR(255) UNIQUE,
   transaction_code VARCHAR(255),
   transaction_type VARCHAR(255),
   total_amount DOUBLE PRECISION NOT NULL DEFAULT 0,
   user_id INT,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);