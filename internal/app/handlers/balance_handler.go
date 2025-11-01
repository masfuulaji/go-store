package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/store/internal/app/repositories"
	"github.com/masfuulaji/store/internal/utils"
)

type BalanceHandlerImpl struct {
	balanceRepository repositories.BalanceRepository
}

func NewBalanceHandler(db *sqlx.DB) *BalanceHandlerImpl {
	return &BalanceHandlerImpl{balanceRepository: repositories.NewBalanceRepository(db)}
}

func (u *BalanceHandlerImpl) GetBalances(w http.ResponseWriter, r *http.Request) {
	user_id := chi.URLParam(r, "id")
	balance, err := u.balanceRepository.GetBalanceByUserId(user_id)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"status":  0,
		"message": "Sukses",
		"data":    balance,
	})
}
