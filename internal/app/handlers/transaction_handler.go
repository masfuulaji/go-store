package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/store/internal/app/models"
	"github.com/masfuulaji/store/internal/app/repositories"
	"github.com/masfuulaji/store/internal/utils"
)

type TransactionHandlerImpl struct {
	db                    *sqlx.DB
	transactionRepository repositories.TransactionRepository
	serviceRepository     repositories.ServiceRepository
	balanceRepository     repositories.BalanceRepository
}

func NewTransactionHandler(db *sqlx.DB) *TransactionHandlerImpl {
	return &TransactionHandlerImpl{
		db:                    db,
		transactionRepository: repositories.NewTransactionRepository(db),
		serviceRepository:     repositories.NewServiceRepository(db),
		balanceRepository:     repositories.NewBalanceRepository(db),
	}
}

func (u *TransactionHandlerImpl) GetTransactions(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.ExtractUserIDFromJWT(r)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]any{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	offsetStr := r.URL.Query().Get("offset")
	limitStr := r.URL.Query().Get("limit")
	offset := 0
	limit := 0

	if offsetStr != "" {
		if p, err := strconv.Atoi(offsetStr); err == nil && p > 0 {
			offset = p
		}
	}

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}
	transactions, err := u.transactionRepository.GetTransactionsByUserId(userID, limit, offset)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]any{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"status":  0,
		"message": "Sukses",
		"data": map[string]any{
			"records": transactions,
			"offset":  offset,
			"limit":   limit,
		},
	})
}

func (u *TransactionHandlerImpl) Transaction(w http.ResponseWriter, r *http.Request) {
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
		ServiceCode string `json:"service_code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	service, err := u.serviceRepository.GetServiceByCode(req.ServiceCode)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": "Service ataus Layanan tidak ditemukan",
			"data":    nil,
		})
		return
	}

	balance, err := u.balanceRepository.GetBalanceByUserId(userID)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	if service.ServiceTariff > balance.Amount {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]interface{}{
			"status":  1,
			"message": "Saldo tidak cukup",
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
	newAmount := balance.Amount - service.ServiceTariff
	balance, err = u.balanceRepository.UpdateBalanceByUserId(tx, newAmount, userID)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	transaction := models.Transaction{
		InvoiceNumber:   fmt.Sprintf("INV-%d", time.Now().UnixNano()),
		TransactionCode: req.ServiceCode,
		TotalAmount:     service.ServiceTariff,
		UserId:          userID,
	}
	resTransaction, err := u.transactionRepository.CreateTransaction(tx, transaction)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	tx.Commit()
	type TransactionResponse struct {
		InvoiceNumber   string    `json:"invoice_number"`
		ServiceCode     string    `json:"service_code"`
		ServiceName     string    `json:"service_name"`
		TransactionType string    `json:"transaction_type"`
		TotalAmount     float64   `json:"total_amount"`
		CreatedAt       time.Time `json:"created_at"`
	}
	transactionResponse := TransactionResponse{
		InvoiceNumber:   resTransaction.InvoiceNumber,
		ServiceCode:     service.ServiceCode,
		ServiceName:     service.ServiceName,
		TransactionType: resTransaction.TransactionType,
		TotalAmount:     resTransaction.TotalAmount,
		CreatedAt:       resTransaction.CreatedAt,
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"status":  0,
		"message": "Transaksi berhasil",
		"data":    transactionResponse,
	})
}
