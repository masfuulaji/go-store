package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/store/internal/app/repositories"
	"github.com/masfuulaji/store/internal/utils"
)

type ServiceHandlerImpl struct {
	serviceRepository repositories.ServiceRepository
}

func NewServiceHandler(db *sqlx.DB) *ServiceHandlerImpl {
	return &ServiceHandlerImpl{serviceRepository: repositories.NewServiceRepository(db)}
}

func (u *ServiceHandlerImpl) GetServices(w http.ResponseWriter, r *http.Request) {
	services, err := u.serviceRepository.GetServices()
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
		"data":    services,
	})
}
