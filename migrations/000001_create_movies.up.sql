CREATE TABLE IF NOT EXISTS movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    price REAL,
    start_date TIMESTAMP,
    end_date TIMESTAMP
);