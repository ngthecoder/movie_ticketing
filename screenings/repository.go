package screenings

import "github.com/jmoiron/sqlx"

type ScreeningRepository struct {
	db *sqlx.DB
}

func NewScreeningRepository(db *sqlx.DB) *ScreeningRepository {
	return &ScreeningRepository{
		db: db,
	}
}

func (r *ScreeningRepository) GetAll() ([]Screening, error) {
	screenings := []Screening{}
	err := r.db.Select(&screenings, "SELECT * FROM screenings")
	if err != nil {
		return nil, err
	}

	return screenings, nil
}

func (r *ScreeningRepository) GetById(id int) (Screening, error) {
	screening := Screening{}
	err := r.db.Get(&screening, "SELECT * FROM screenings WHERE id=$1", id)
	if err != nil {
		return Screening{}, err
	}

	return screening, nil
}
