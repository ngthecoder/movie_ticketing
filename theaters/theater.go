package theaters

type Theater struct {
	Id         int    `db:"id"`
	Name       string `db:"name"`
	Location   string `db:"location"`
	TotalSeats int    `db:"total_seats"`
}
