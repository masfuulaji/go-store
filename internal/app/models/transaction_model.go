package models

import (
	"time"

	"github.com/lib/pq"
)

type Transaction struct {
	CreatedAt       time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time   `db:"updated_at" json:"updated_at"`
	DeletedAt       pq.NullTime `db:"deleted_at" json:"deleted_at"`
	InvoiceNumber   string      `db:"invoice_number" json:"invoice_number"`
	TransactionType string      `db:"transaction_type" json:"transaction_type"`
	TotalAmount     float64     `db:"total_amount" json:"total_amount"`
	UserId          string      `db:"user_id" json:"user_id"`
	Id              int         `json:"id"`
}
