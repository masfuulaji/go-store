package models

import (
	"time"

	"github.com/lib/pq"
)

type Banner struct {
	CreatedAt   time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time   `db:"updated_at" json:"updated_at"`
	DeletedAt   pq.NullTime `db:"deleted_at" json:"deleted_at"`
	BannerName  string      `db:"banner_name" json:"banner_name"`
	BannerImage string      `db:"banner_image" json:"banner_image"`
	Description string      `json:"description"`
	Id          int         `json:"id"`
}
