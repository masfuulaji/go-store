package models

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type User struct {
	CreatedAt    time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at" json:"updated_at"`
	DeletedAt    pq.NullTime    `db:"deleted_at" json:"deleted_at"`
	FirstName    string         `db:"first_name" json:"first_name"`
	LastName     string         `db:"last_name" json:"last_name"`
	ProfileImage sql.NullString `db:"profile_image" json:"profile_image,omitempty"`
	Email        string         `json:"email"`
	Password     string         `json:"password"`
	Id           int            `json:"id"`
}
