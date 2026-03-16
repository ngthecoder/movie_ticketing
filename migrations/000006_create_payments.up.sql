CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    booking_id INT REFERENCES bookings(id) UNIQUE,
    amount REAL,
    paid_at TIMESTAMP DEFAULT NOW()
);