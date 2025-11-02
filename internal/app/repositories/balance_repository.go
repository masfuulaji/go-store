package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/store/internal/app/models"
)

type BalanceRepository interface {
	CreateBalance(tx *sqlx.Tx, balance float64, userID int) (models.Balance, error)
	GetBalanceByUserId(user_id int) (models.Balance, error)
	UpdateBalanceByUserId(tx *sqlx.Tx, balance float64, id int) (models.Balance, error)
}

type BalanceRepositoryImpl struct {
	db *sqlx.DB
}

func NewBalanceRepository(db *sqlx.DB) BalanceRepository {
	return &BalanceRepositoryImpl{db: db}
}

func (u *BalanceRepositoryImpl) CreateBalance(tx *sqlx.Tx, balance float64, userID int) (models.Balance, error) {
	var resBalance models.Balance
	query := "INSERT INTO balances (user_id, amount, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, user_id, amount, created_at, updated_at"
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	// _, err := u.db.Exec(query, userID, balance, createdAt, updatedAt)
	err := tx.Get(&resBalance, query, userID, balance, createdAt, updatedAt)
	if err != nil {
		return resBalance, err
	}
	return resBalance, nil
}

func (u *BalanceRepositoryImpl) GetBalanceByUserId(userID int) (models.Balance, error) {
	var balance models.Balance
	query := "SELECT * FROM balances WHERE user_id = $1 AND deleted_at IS NULL"
	err := u.db.Get(&balance, query, userID)
	if err != nil {
		return balance, err
	}
	return balance, nil
}

func (u *BalanceRepositoryImpl) UpdateBalanceByUserId(tx *sqlx.Tx, balance float64, userID int) (models.Balance, error) {
	var resBalance models.Balance
	query := `
		UPDATE balances
		SET amount = $1, updated_at = $2
		WHERE user_id = $3
		RETURNING *;
	`

	updatedAt := time.Now()
	err := tx.Get(&resBalance, query, balance, updatedAt, userID)
	if err != nil {
		return resBalance, err
	}

	return resBalance, nil
}
