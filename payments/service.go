package payments

import (
	"github.com/ngthecoder/movie_ticketing/bookings"
	"github.com/ngthecoder/movie_ticketing/movies"
	"github.com/ngthecoder/movie_ticketing/screenings"
)

type PaymentService struct {
	r          *PaymentRepository
	bookingS   *bookings.BookingService
	screeningS *screenings.ScreeningService
	movieS     *movies.MovieService
}

func NewPaymentService(r *PaymentRepository, bookingS *bookings.BookingService, screeningS *screenings.ScreeningService, movieS *movies.MovieService) *PaymentService {
	return &PaymentService{
		r:          r,
		bookingS:   bookingS,
		screeningS: screeningS,
		movieS:     movieS,
	}
}

func (s *PaymentService) Pay(bookingId int) (Payment, error) {
	booking, err := s.bookingS.GetById(bookingId)
	if err != nil {
		return Payment{}, err
	}

	screening, err := s.screeningS.GetById(booking.ScreeningId)
	if err != nil {
		return Payment{}, err
	}

	movie, err := s.movieS.GetById(screening.MovieId)
	if err != nil {
		return Payment{}, err
	}

	amount := movie.Price * float32(booking.NumTickets)
	payment, err := s.r.Pay(bookingId, amount)
	if err != nil {
		return Payment{}, err
	}

	return payment, nil
}
