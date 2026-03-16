CREATE TABLE IF NOT EXISTS screenings (
    id SERIAL PRIMARY KEY,
    movie_id INT REFERENCES movies(id),
    theater_id INT REFERENCES theaters(id),
    starts_at TIMESTAMP,
    available_seats INT,
    UNIQUE(movie_id, theater_id, starts_at)
);