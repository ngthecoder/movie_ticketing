package theaters

import (
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("theaters not found")

type TheaterService struct {
	r *TheaterRepository
}

func NewTheaterService(r *TheaterRepository) *TheaterService {
	return &TheaterService{
		r: r,
	}
}

func (s *TheaterService) GetAll() ([]Theater, error) {
	theaters, err := s.r.GetAll()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return theaters, nil
}

func (s *TheaterService) GetById(id int) (Theater, error) {
	theater, err := s.r.GetById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Theater{}, ErrNotFound
		}

		return Theater{}, err
	}

	return theater, nil
}
