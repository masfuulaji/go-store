package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/store/internal/app/models"
)

type ServiceRepository interface {
	GetServices() ([]models.Service, error)
	GetServiceByCode(string) (models.Service, error)
}

type ServiceRepositoryImpl struct {
	db *sqlx.DB
}

func NewServiceRepository(db *sqlx.DB) ServiceRepository {
	return &ServiceRepositoryImpl{db: db}
}

func (u ServiceRepositoryImpl) GetServices() ([]models.Service, error) {
	query := "SELECT * FROM services WHERE deleted_at IS NULL"
	var services []models.Service
	err := u.db.Select(&services, query)
	if err != nil {
		return services, err
	}
	return services, nil
}

func (u ServiceRepositoryImpl) GetServiceByCode(serviceCode string) (models.Service, error) {
	var service models.Service
	query := "SELECT * FROM services WHERE service_code = $1 AND deleted_at IS NULL"
	err := u.db.Get(&service, query, serviceCode)
	if err != nil {
		return service, err
	}
	return service, nil
}
