package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/store/internal/app/repositories"
	"github.com/masfuulaji/store/internal/utils"
)

type TransactionHandlerImpl struct {
	transactionRepository repositories.TransactionRepository
}

func NewTransactionHandler(db *sqlx.DB) *TransactionHandlerImpl {
	return &TransactionHandlerImpl{transactionRepository: repositories.NewTransactionRepository(db)}
}

func (u *TransactionHandlerImpl) GetTransactions(w http.ResponseWriter, r *http.Request) {
	user_id := chi.URLParam(r, "id")
	transaction, err := u.transactionRepository.GetTransactionByUserId(user_id)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	json.NewEncoder(w).Encode(transaction)
}
