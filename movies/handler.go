package movies

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	service *MovieService
}

func NewMovieHandler(service *MovieService) *MovieHandler {
	movieHandler := &MovieHandler{
		service: service,
	}

	return movieHandler
}

func (h *MovieHandler) GetAllHandler(c *gin.Context) {
	movies, err := h.service.GetAll()
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "movie not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"movies": movies,
	})
}

func (h *MovieHandler) GetByIdHandler(c *gin.Context) {
	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID format not compatible",
		})
		return
	}

	movie, err := h.service.GetById(id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "movie not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"movie": movie,
	})
}
