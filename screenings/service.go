package screenings

import (
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("screening not found")

type ScreeningService struct {
	r *ScreeningRepository
}

func NewScreeningService(r *ScreeningRepository) *ScreeningService {
	return &ScreeningService{
		r: r,
	}
}

func (s *ScreeningService) GetAll() ([]Screening, error) {
	screenings, err := s.r.GetAll()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return screenings, nil
}

func (s *ScreeningService) GetById(id int) (Screening, error) {
	screening, err := s.r.GetById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Screening{}, ErrNotFound
		}

		return Screening{}, err
	}

	return screening, nil
}
