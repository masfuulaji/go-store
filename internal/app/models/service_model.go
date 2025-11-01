package models

import (
	"time"

	"github.com/lib/pq"
)

type Service struct {
	CreatedAt   time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time   `db:"updated_at" json:"updated_at"`
	DeletedAt   pq.NullTime `db:"deleted_at" json:"deleted_at"`
	ServiceCode string      `db:"service_code" json:"service_code"`
	ServiceName string      `db:"sservice_name" json:"sservice_name"`
	ServiceIcon string      `db:"service_icon" json:"service_icon"`
	Id          int         `json:"id"`
}
