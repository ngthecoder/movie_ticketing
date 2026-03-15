package movies

import "time"

type Movie struct {
	Id          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Price       float32   `db:"price"`
	StartDate   time.Time `db:"start_date"`
	EndDate     time.Time `db:"end_date"`
}
