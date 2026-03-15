package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngthecoder/movie_ticketing/db"
	"github.com/ngthecoder/movie_ticketing/movies"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer db.Close()

	movieRepository := movies.NewMovieRepository(db)
	movieService := movies.NewMovieService(movieRepository)
	movieHandler := movies.NewMovieHandler(movieService)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/movies", movieHandler.GetAllHandler)
	r.GET("/movies/:id", movieHandler.GetByIdHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
