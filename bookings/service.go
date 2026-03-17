package bookings

import (
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("booking not found")
var ErrScreeningNotFound = errors.New("screening not found")

type BookingService struct {
	r *BookingRepository
}

func NewBookingService(r *BookingRepository) *BookingService {
	return &BookingService{
		r: r,
	}
}

func (s *BookingService) Create(userId, screeningId, numTickets int) (Booking, error) {
	booking, err := s.r.Create(userId, screeningId, numTickets)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Booking{}, ErrScreeningNotFound
		}
		return Booking{}, err
	}

	return booking, nil
}

func (s *BookingService) GetById(id int) (Booking, error) {
	booking, err := s.r.GetById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Booking{}, ErrNotFound
		}

		return Booking{}, err
	}

	return booking, nil
}
