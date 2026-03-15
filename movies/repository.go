package movies

import "github.com/jmoiron/sqlx"

type MovieRepository struct {
	db *sqlx.DB
}

func NewMovieRepository(db *sqlx.DB) *MovieRepository {
	movieRepository := &MovieRepository{
		db: db,
	}

	return movieRepository
}

func (r *MovieRepository) GetAll() ([]Movie, error) {
	movies := []Movie{}
	err := r.db.Select(&movies, "SELECT * FROM movies")
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *MovieRepository) GetById(id int) (Movie, error) {
	movie := Movie{}
	err := r.db.Get(&movie, "SELECT * FROM movies WHERE id=$1", id)
	if err != nil {
		return Movie{}, err
	}

	return movie, nil
}
