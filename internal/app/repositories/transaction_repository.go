package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/store/internal/app/models"
)

type TransactionRepository interface {
	CreateTransaction(transaction models.Transaction) error
	GetTransactionByUserId(userID string) (models.Transaction, error)
	UpdateTransaction(transaction models.Transaction, id string) error
}

type TransactionRepositoryImpl struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &TransactionRepositoryImpl{db: db}
}

func (u TransactionRepositoryImpl) CreateTransaction(transaction models.Transaction) error {
	query := "INSERT INTO transactions (user_id, invoice_number, transaction_type, total_amount, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	_, err := u.db.Exec(query, transaction.UserId, transaction.InvoiceNumber, transaction.TransactionType, transaction.TotalAmount, createdAt, updatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (u TransactionRepositoryImpl) GetTransactionByUserId(user_id string) (models.Transaction, error) {
	var transaction models.Transaction
	query := "SELECT * FROM transactions WHERE user_id = $1 AND deleted_at IS NULL"
	err := u.db.Get(&transaction, query, user_id)
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
