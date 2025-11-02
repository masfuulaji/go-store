package models

import (
	"time"

	"github.com/lib/pq"
)

type Service struct {
	CreatedAt     time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time   `db:"updated_at" json:"updated_at"`
	DeletedAt     pq.NullTime `db:"deleted_at" json:"deleted_at"`
	ServiceCode   string      `db:"service_code" json:"service_code"`
	ServiceName   string      `db:"service_name" json:"service_name"`
	ServiceIcon   string      `db:"service_icon" json:"service_icon"`
	ServiceTariff float64     `db:"service_tariff" json:"service_tariff"`
	Id            int         `json:"id"`
}
