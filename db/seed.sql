TRUNCATE TABLE screenings, theaters, movies, users RESTART IDENTITY CASCADE;

INSERT INTO movies (title, description, price, start_date, end_date) VALUES
('Inception', 'A mind-bending thriller', 15.00, '2025-01-01', '2025-12-31'),
('Interstellar', 'A space odyssey', 12.00, '2025-01-01', '2025-12-31');

INSERT INTO theaters (name, location, total_seats) VALUES
('Theater A', 'Tokyo', 100),
('Theater B', 'Osaka', 80);

INSERT INTO users (email, password_hash) VALUES
('micheal.jackson@gmail.com', '$2y$10$HUHIWiimJ5gOT3eJJFZfc.xztnKdiGXEqoFrjZVWPRMFjMYdnw0GG%'),
('steve.jobs@gmail.com', '$2y$10$HUHIWiimJ5gOT3eJJFZfc.xztnKdiGXEqoFrjZVWPRMFjMYdnw0GG%'),
('naoki.goto@gmail.com', '$2y$10$HUHIWiimJ5gOT3eJJFZfc.xztnKdiGXEqoFrjZVWPRMFjMYdnw0GG%');

INSERT INTO screenings (movie_id, theater_id, starts_at, available_seats) VALUES
(1, 1, '2025-06-01 10:00:00', 100),
(1, 2, '2025-06-01 14:00:00', 80),
(2, 1, '2025-06-02 18:00:00', 100);