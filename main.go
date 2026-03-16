package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngthecoder/movie_ticketing/db"
	"github.com/ngthecoder/movie_ticketing/movies"
	"github.com/ngthecoder/movie_ticketing/theaters"
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

	theaterRepository := theaters.NewTheaterRepository(db)
	theaterService := theaters.NewTheaterService(theaterRepository)
	theaterHandler := theaters.NewTheaterHandler(theaterService)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/movies", movieHandler.GetAllHandler)
	r.GET("/movies/:id", movieHandler.GetByIdHandler)

	r.GET("/theaters", theaterHandler.GetAllHandler)
	r.GET("/theaters/:id", theaterHandler.GetByIdHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
