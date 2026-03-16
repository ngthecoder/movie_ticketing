package payments

import "time"

type Payment struct {
	Id        int       `db:"id"`
	BookingId int       `db:"booking_id"`
	Amount    float32   `db:"amount"`
	PaidAt    time.Time `db:"paid_at"`
}
