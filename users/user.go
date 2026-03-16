package users

import "time"

type User struct {
	Id           int       `db:"id"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	CreateAt     time.Time `db:"created_at"`
}
