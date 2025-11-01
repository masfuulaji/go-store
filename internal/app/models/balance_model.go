package models

import (
	"time"

	"github.com/lib/pq"
)

type Balance struct {
	CreatedAt time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt time.Time   `db:"updated_at" json:"updated_at"`
	DeletedAt pq.NullTime `db:"deleted_at" json:"deleted_at"`
	UserId    string      `db:"user_id" json:"user_id"`
	Amount    float64     `json:"amount"`
	Id        int         `json:"id"`
}
