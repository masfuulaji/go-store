package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/store/internal/app/models"
	"github.com/masfuulaji/store/internal/app/repositories"
	"github.com/masfuulaji/store/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserHandlerImpl struct {
	userRepository repositories.UserRepository
}

func NewUserHandler(db *sqlx.DB) *UserHandlerImpl {
	return &UserHandlerImpl{userRepository: repositories.NewUserRepository(db)}
}

func (u *UserHandlerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)
	err = u.userRepository.CreateUser(user)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	utils.RespondWithJSON(w, 200, map[string]string{"message": "User created successfully"})
}

func (u *UserHandlerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	err = u.userRepository.UpdateUser(user, id)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	utils.RespondWithJSON(w, 200, map[string]string{"message": "User created successfully"})
}

func (u *UserHandlerImpl) UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Image is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	imageURL, err := utils.SaveUploadedFile(file, handler, "uploads")
	if err != nil {
		http.Error(w, "Failed to save image: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = u.userRepository.UpdateUserProfile(imageURL, id)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	utils.RespondWithJSON(w, 200, map[string]string{"message": "User created successfully"})
}

func (u *UserHandlerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := u.userRepository.DeleteUser(id)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	utils.RespondWithJSON(w, 200, map[string]string{"message": "User created successfully"})
}

func (u *UserHandlerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, err := u.userRepository.GetUser(id)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (u *UserHandlerImpl) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.userRepository.GetUsers()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	json.NewEncoder(w).Encode(users)
}
