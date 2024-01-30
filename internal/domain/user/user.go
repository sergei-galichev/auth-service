package user

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64        `db:"id"`
	UUID      string       `db:"uuid"`
	Email     string       `db:"email"`
	PassHash  []byte       `db:"pass_hash"`
	Role      string       `db:"role"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	LoggedIn  sql.NullTime `db:"logged_in"`
	LoggedOut sql.NullTime `db:"logged_out"`
}

var Roles = map[int32]string{
	0: "EMPLOYEE",
	1: "MANAGER",
	2: "ADMIN",
}

func NewUser(
	uuid string,
	email string,
	passHash []byte,
	role string,
) User {
	return User{
		UUID:      uuid,
		Email:     email,
		PassHash:  passHash,
		Role:      role,
		CreatedAt: time.Now(),
	}
}
