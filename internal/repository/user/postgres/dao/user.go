package dao

import (
	"database/sql"
	"github.com/upper/db/v4"
	"time"
)

var (
	_ = db.Record(&UserDAO{})
)

type UserDAO struct {
	ID        int64        `db:"id,omitempty"`
	UUID      string       `db:"uuid"`
	Email     string       `db:"email"`
	PassHash  []byte       `db:"pass_hash"`
	Role      string       `db:"role"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	LoggedIn  sql.NullTime `db:"logged_in"`
	LoggedOut sql.NullTime `db:"logged_out"`
}

func (*UserDAO) Store(session db.Session) db.Store {
	return session.Collection("users")
}
