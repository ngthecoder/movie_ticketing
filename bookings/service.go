package bookings

import (
	"errors"

	"github.com/ngthecoder/movie_ticketing/screenings"
)

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
