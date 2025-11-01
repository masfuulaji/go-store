package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/store/internal/app/models"
)

var balance models.Balance

type BalanceRepository interface {
	CreateBalance(balance models.Balance) error
	GetBalanceByUserId(user_id string) (models.Balance, error)
	UpdateBalanceByUserId(balance models.Balance, id string) error
}

type BalanceRepositoryImpl struct {
	db *sqlx.DB
}

func NewBalanceRepository(db *sqlx.DB) BalanceRepository {
	return &BalanceRepositoryImpl{db: db}
}

func (u *BalanceRepositoryImpl) CreateBalance(balance models.Balance) error {
	query := "INSERT INTO balances (user_id, amount, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	_, err := u.db.Exec(query, balance.UserId, balance.Amount, createdAt, updatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (u *BalanceRepositoryImpl) GetBalanceByUserId(user_id string) (models.Balance, error) {
	query := "SELECT * FROM balances WHERE user_id = $1 AND deleted_at IS NULL"
	err := u.db.Get(&balance, query, user_id)
	if err != nil {
		return balance, err
	}
	return balance, nil
}

func (u *BalanceRepositoryImpl) UpdateBalanceByUserId(balance models.Balance, user_id string) error {
	query := "UPDATE balances SET amount = $1, updated_at = $2 WHERE user_id = $3"
	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	_, err := u.db.Exec(query, balance.Amount, updatedAt, user_id)
	if err != nil {
		return err
	}
	return nil
}
