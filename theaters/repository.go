package theaters

import "github.com/jmoiron/sqlx"

type TheaterRepository struct {
	db *sqlx.DB
}

func NewTheaterRepository(db *sqlx.DB) *TheaterRepository {
	return &TheaterRepository{
		db: db,
	}
}

func (r *TheaterRepository) GetAll() ([]Theater, error) {
	theaters := []Theater{}
	err := r.db.Select(&theaters, "SELECT * FROM theaters")
	if err != nil {
		return nil, err
	}

	return theaters, nil
}

func (r *TheaterRepository) GetById(id int) (Theater, error) {
	theater := Theater{}
	err := r.db.Get(&theater, "SELECT * FROM theaters WHERE id=$1", id)
	if err != nil {
		return Theater{}, err
	}

	return theater, nil
}
