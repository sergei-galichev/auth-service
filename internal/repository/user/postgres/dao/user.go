package dao

import (
	"database/sql"
	"time"
)

type UserDAO struct {
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
