package bookings

import "time"

type Booking struct {
	Id          int       `db:"id"`
	UserId      int       `db:"user_id"`
	ScreeningId int       `db:"screening_id"`
	NumTickets  int       `db:"num_tickets"`
	Status      string    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
}
