package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/store/internal/app/repositories"
	"github.com/masfuulaji/store/internal/utils"
)

type BalanceHandlerImpl struct {
	db                *sqlx.DB
	balanceRepository repositories.BalanceRepository
}

func NewBalanceHandler(db *sqlx.DB) *BalanceHandlerImpl {
	return &BalanceHandlerImpl{
		db:                db,
		balanceRepository: repositories.NewBalanceRepository(db),
	}
}

func (u *BalanceHandlerImpl) GetBalance(w http.ResponseWriter, r *http.Request) {
	// id := chi.URLParam(r, "id")
	userID, err := utils.ExtractUserIDFromJWT(r)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	balance, err := u.balanceRepository.GetBalanceByUserId(userID)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": "Balance tidak ditemukan",
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

func (u *BalanceHandlerImpl) TopUp(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.ExtractUserIDFromJWT(r)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	var req struct {
		TopUpAmount float64 `json:"top_up_amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	if req.TopUpAmount <= 0 {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]interface{}{
			"status":  1,
			"message": "Top up amount must be greater than zero",
			"data":    nil,
		})
		return
	}

	tx, err := u.db.Beginx()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	defer tx.Rollback()

	balance, err := u.balanceRepository.GetBalanceByUserId(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			balance, err = u.balanceRepository.CreateBalance(tx, req.TopUpAmount, userID)
			if err != nil {
				utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
					"status":  1,
					"message": err.Error(),
					"data":    nil,
				})
				return
			}
			if err := tx.Commit(); err != nil {
				utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
					"status":  1,
					"message": "Failed to commit transaction",
					"data":    nil,
				})
				return
			}

			utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
				"status":  0,
				"message": "Balance created and top-up successful",
				"data":    balance,
			})
			return
		}

		// Other unexpected errors
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	newAmount := balance.Amount + req.TopUpAmount
	balance, err = u.balanceRepository.UpdateBalanceByUserId(tx, newAmount, userID)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	if err := tx.Commit(); err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": "Failed to commit transaction",
			"data":    nil,
		})
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"status":  0,
		"message": "Top Up Balance berhasil",
		"data":    balance,
	})
}
