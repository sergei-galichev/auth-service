package user

import (
	"database/sql"
	"regexp"
	"strings"
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

var regexpEmail = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

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

func IsValidEmail(email string) bool {
	return regexpEmail.MatchString(strings.ToLower(email))
}
