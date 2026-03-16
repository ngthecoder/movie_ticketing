package users

import "github.com/jmoiron/sqlx"

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(email string, passwordHash string) (User, error) {
	user := User{}
	err := r.db.Get(&user, "INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING *", email, passwordHash)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) GetByEmail(email string) (User, error) {
	user := User{}
	err := r.db.Get(&user, "SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		return user, err
	}

	return user, nil
}
