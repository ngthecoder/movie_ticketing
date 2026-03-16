package bookings

import (
	"database/sql"
	"errors"

	"github.com/ngthecoder/movie_ticketing/screenings"
)

var ErrNotFound = errors.New("booking not found")
var ErrNotEnoughSeats = errors.New("not enough seats available")

type BookingService struct {
	r          *BookingRepository
	screeningS *screenings.ScreeningService
}

func NewBookingService(r *BookingRepository, screeningS *screenings.ScreeningService) *BookingService {
	return &BookingService{
		r:          r,
		screeningS: screeningS,
	}
}

func (s *BookingService) Create(userId, screeningId, numTickets int) (Booking, error) {
	screening, err := s.screeningS.GetById(screeningId)
	if err != nil {
		if errors.Is(err, screenings.ErrNotFound) {
			return Booking{}, screenings.ErrNotFound
		}

		return Booking{}, err
	}
	if numTickets > screening.AvailableSeats {
		return Booking{}, ErrNotEnoughSeats
	}

	booking, err := s.r.Create(userId, screeningId, numTickets)
	if err != nil {
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
