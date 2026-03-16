package payments

import "github.com/jmoiron/sqlx"

type PaymentRepository struct {
	db *sqlx.DB
}

func NewPaymentRepository(db *sqlx.DB) *PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

func (r *PaymentRepository) Pay(bookingId int, amount float32) (Payment, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return Payment{}, err
	}
	defer tx.Rollback()

	payment := Payment{}
	err = tx.Get(&payment, "INSERT INTO payments (booking_id, amount) VALUES ($1, $2) RETURNING *", bookingId, amount)
	if err != nil {
		return Payment{}, err
	}

	status := "confirmed"
	_, err = tx.Exec("UPDATE bookings SET status=$2 WHERE id=$1", bookingId, status)
	if err != nil {
		return Payment{}, err
	}

	tx.Commit()
	return payment, nil
}
