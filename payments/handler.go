package payments

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngthecoder/movie_ticketing/bookings"
	"github.com/ngthecoder/movie_ticketing/movies"
	"github.com/ngthecoder/movie_ticketing/screenings"
)

type PaymentHandler struct {
	s *PaymentService
}

func NewPaymentHandler(s *PaymentService) *PaymentHandler {
	return &PaymentHandler{
		s: s,
	}
}

func (h *PaymentHandler) PayHandler(c *gin.Context) {
	var req struct {
		BookingId int `json:"booking_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	payment, err := h.s.Pay(req.BookingId)
	if err != nil {
		if errors.Is(err, bookings.ErrNotFound) || errors.Is(err, screenings.ErrNotFound) || errors.Is(err, movies.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "resource not found",
			})
			return
		}

		log.Printf("internal error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payment": payment,
	})
}
