package bookings

import "github.com/jmoiron/sqlx"

type BookingRepository struct {
	db *sqlx.DB
}

func NewBookingRepository(db *sqlx.DB) *BookingRepository {
	return &BookingRepository{
		db: db,
	}
}

func (r *BookingRepository) Create(userId, screeningId, numTickets int) (Booking, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return Booking{}, err
	}
	defer tx.Rollback()

	booking := Booking{}
	status := "pending"
	err = tx.Get(&booking, "INSERT INTO bookings (user_id, screening_id, num_tickets, status) VALUES ($1, $2, $3, $4) RETURNING *", userId, screeningId, numTickets, status)
	if err != nil {
		return Booking{}, err
	}

	_, err = tx.Exec("UPDATE screenings SET available_seats=available_seats-$2 WHERE id=$1", screeningId, numTickets)
	if err != nil {
		return Booking{}, err
	}

	tx.Commit()
	return booking, nil
}

func (r *BookingRepository) GetById(id int) (Booking, error) {
	booking := Booking{}
	err := r.db.Get(&booking, "SELECT * FROM bookings WHERE id=$1", id)
	if err != nil {
		return Booking{}, err
	}

	return booking, nil
}
