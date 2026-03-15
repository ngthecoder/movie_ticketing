package movies

import (
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("movie not found")

type MovieService struct {
	repository *MovieRepository
}

func NewMovieService(repository *MovieRepository) *MovieService {
	movieService := &MovieService{
		repository: repository,
	}

	return movieService
}

func (s *MovieService) GetAll() ([]Movie, error) {
	movies, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (s *MovieService) GetById(id int) (Movie, error) {
	movie, err := s.repository.GetById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Movie{}, ErrNotFound
		}
		return Movie{}, err
	}

	return movie, nil
}
