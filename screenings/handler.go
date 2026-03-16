package screenings

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ScreeningHandler struct {
	s *ScreeningService
}

func NewScreeningHandler(s *ScreeningService) *ScreeningHandler {
	return &ScreeningHandler{
		s: s,
	}
}

func (h *ScreeningHandler) GetAllHandler(c *gin.Context) {
	screenings, err := h.s.GetAll()
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "screening not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"screenings": screenings,
	})
}

func (h *ScreeningHandler) GetByIdHandler(c *gin.Context) {
	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID format not compatible",
		})
		return
	}

	screening, err := h.s.GetById(id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "screening not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"screening": screening,
	})
}
