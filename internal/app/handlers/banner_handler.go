package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/store/internal/app/repositories"
	"github.com/masfuulaji/store/internal/utils"
)

type BannerHandlerImpl struct {
	bannerRepository repositories.BannerRepository
}

func NewBannerHandler(db *sqlx.DB) *BannerHandlerImpl {
	return &BannerHandlerImpl{bannerRepository: repositories.NewBannerRepository(db)}
}

func (u *BannerHandlerImpl) GetBanners(w http.ResponseWriter, r *http.Request) {
	banners, err := u.bannerRepository.GetBanners()
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
		"data":    banners,
	})
}
