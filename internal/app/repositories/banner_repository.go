package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/store/internal/app/models"
)

type BannerRepository interface {
	GetBanners() ([]models.Banner, error)
}

type BannerRepositoryImpl struct {
	db *sqlx.DB
}

func NewBannerRepository(db *sqlx.DB) BannerRepository {
	return &BannerRepositoryImpl{db: db}
}

func (u BannerRepositoryImpl) GetBanners() ([]models.Banner, error) {
	query := "SELECT * FROM banners WHERE deleted_at IS NULL"
	var banners []models.Banner
	err := u.db.Select(&banners, query)
	if err != nil {
		return banners, err
	}
	return banners, nil
}
