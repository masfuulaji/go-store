package models

import (
	"time"

	"github.com/lib/pq"
)

type User struct {
	CreatedAt time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt time.Time   `db:"updated_at" json:"updated_at"`
	DeletedAt pq.NullTime `db:"deleted_at" json:"deleted_at"`
	FirstName string      `db:"first_name" json:"first_name"`
	LastName  string      `db:"last_name" json:"last_name"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	Id        int         `json:"id"`
}
