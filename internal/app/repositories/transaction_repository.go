package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/store/internal/app/models"
)

type TransactionRepository interface {
	CreateTransaction(tx *sqlx.Tx, transaction models.Transaction) (models.Transaction, error)
	GetTransactionsByUserId(userID, limit, offset int) (models.Transaction, error)
	UpdateTransaction(transaction models.Transaction, id string) error
}

type TransactionRepositoryImpl struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &TransactionRepositoryImpl{db: db}
}

func (u TransactionRepositoryImpl) CreateTransaction(tx *sqlx.Tx, transaction models.Transaction) (models.Transaction, error) {
	var resTransaction models.Transaction
	query := `INSERT INTO transactions (user_id, invoice_number, transaction_code, transaction_type, total_amount, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)
				RETURNING id, invoice_number, transaction_code, transaction_type, total_amount, created_at, updated_at`
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	err := tx.Get(&resTransaction, query, transaction.UserId, transaction.InvoiceNumber, transaction.TransactionCode, "Payment", transaction.TotalAmount, createdAt, updatedAt)
	if err != nil {
		return resTransaction, err
	}
	return resTransaction, nil
}

func (u TransactionRepositoryImpl) GetTransactionsByUserId(userID, limit, offset int) (models.Transaction, error) {
	var transaction models.Transaction
	var query string
	var args []any
	if limit > 0 {
		query = "SELECT * FROM transactions WHERE user_id = $1 AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $2 OFFSET $3"
		args = []any{userID, limit, offset}
	} else {
		query = "SELECT * FROM transactions WHERE user_id = $1 AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $2"
		args = []any{userID, offset}
	}
	err := u.db.Get(&transaction, query, args...)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (u TransactionRepositoryImpl) UpdateTransaction(transaction models.Transaction, id string) error {
	query := "UPDATE transactions SET total_amount = $1, updated_at = $2 WHERE user_id = $3"
	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	_, err := u.db.Exec(query, transaction.TotalAmount, updatedAt, id)
	if err != nil {
		return err
	}
	return nil
}
