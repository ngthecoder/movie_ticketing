package theaters

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TheaterHandler struct {
	s *TheaterService
}

func NewTheaterHandler(s *TheaterService) *TheaterHandler {
	return &TheaterHandler{
		s: s,
	}
}

func (h *TheaterHandler) GetAllHandler(c *gin.Context) {
	theaters, err := h.s.GetAll()
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "theater not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"theaters": theaters,
	})
}

func (h *TheaterHandler) GetByIdHandler(c *gin.Context) {
	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID format not compatible",
		})
		return
	}

	theater, err := h.s.GetById(id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "theater not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"theater": theater,
	})
}
