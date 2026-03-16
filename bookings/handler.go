package bookings

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngthecoder/movie_ticketing/screenings"
)

type BookingHandler struct {
	s *BookingService
}

func NewBookingHandler(s *BookingService) *BookingHandler {
	return &BookingHandler{
		s: s,
	}
}

func (h *BookingHandler) CreateHandler(c *gin.Context) {
	var req struct {
		UserId      int `json:"user_id"`
		ScreeningId int `json:"screening_id"`
		NumTickets  int `json:"num_tickets"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	booking, err := h.s.Create(req.UserId, req.ScreeningId, req.NumTickets)
	if err != nil {
		if errors.Is(err, ErrNotEnoughSeats) {
			c.JSON(http.StatusConflict, gin.H{
				"error": "not enough seats available",
			})
			return
		} else if errors.Is(err, screenings.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "screening not found",
			})
			return
		}

		log.Printf("internal error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"booking": booking,
	})
}
