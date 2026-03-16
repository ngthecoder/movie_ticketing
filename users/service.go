package users

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

var ErrEmailAlreadyExists = errors.New("email already exists")
var ErrUnauthorized = errors.New("incorrect password")

type UserService struct {
	r *UserRepository
}

func NewUserService(r *UserRepository) *UserService {
	return &UserService{
		r: r,
	}
}

func (s *UserService) CreateUser(email string, password string) (User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	user, err := s.r.CreateUser(email, string(passwordHash))
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return User{}, ErrEmailAlreadyExists
		}

		return User{}, err
	}

	return user, nil
}

func (s *UserService) LoginUser(email string, password string) (User, error) {
	user, err := s.r.GetByEmail(email)
	if err != nil {
		return User{}, ErrUnauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return User{}, ErrUnauthorized
	}

	return user, nil
}
