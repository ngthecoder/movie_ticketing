package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngthecoder/movie_ticketing/bookings"
	"github.com/ngthecoder/movie_ticketing/db"
	"github.com/ngthecoder/movie_ticketing/movies"
	"github.com/ngthecoder/movie_ticketing/screenings"
	"github.com/ngthecoder/movie_ticketing/theaters"
	"github.com/ngthecoder/movie_ticketing/users"
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

	userRepository := users.NewUserRepository(db)
	userService := users.NewUserService(userRepository)
	userHandler := users.NewUserHandler(userService)

	screeningRepository := screenings.NewScreeningRepository(db)
	screeningService := screenings.NewScreeningService(screeningRepository)
	screeningHandler := screenings.NewScreeningHandler(screeningService)

	bookingRepository := bookings.NewBookingRepository(db)
	bookingService := bookings.NewBookingService(bookingRepository, screeningService)
	bookingHandler := bookings.NewBookingHandler(bookingService)

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

	r.POST("/users/register", userHandler.RegisterHandler)
	r.POST("/users/login", userHandler.LoginHandler)

	r.GET("/screenings", screeningHandler.GetAllHandler)
	r.GET("/screenings/:id", screeningHandler.GetByIdHandler)

	r.POST("/bookings", bookingHandler.CreateHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
