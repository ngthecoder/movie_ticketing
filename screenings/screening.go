package screenings

import "time"

type Screening struct {
	Id             int       `db:"id"`
	MovieId        int       `db:"movie_id"`
	TheaterId      int       `db:"theater_id"`
	StartsAt       time.Time `db:"starts_at"`
	AvailableSeats int       `db:"available_seats"`
}
