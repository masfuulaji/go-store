package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/masfuulaji/store/internal/app/models"
	"github.com/masfuulaji/store/internal/app/repositories"
	"github.com/masfuulaji/store/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandlerImpl struct {
	userRepository repositories.UserRepository
}

func NewLoginHandler(db *sqlx.DB) *LoginHandlerImpl {
	return &LoginHandlerImpl{userRepository: repositories.NewUserRepository(db)}
}

type Claims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func (u *LoginHandlerImpl) Login(w http.ResponseWriter, r *http.Request) {
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
	res, err := u.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(user.Password))
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid credentials"})
		return
	}

	expirationTime := time.Now().Add(10 * time.Hour)
	claim := &Claims{
		Id:    res.Id,
		Email: res.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  1,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})

	type TokenResponse struct {
		Token string `json:"token"`
	}
	tokenResponse := TokenResponse{
		Token: tokenString,
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"status":  0,
		"message": "Login Sukses",
		"data":    tokenResponse,
	})
}

func (u *LoginHandlerImpl) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now(),
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"status":  0,
		"message": "Logout Sukses",
		"data":    nil,
	})
}

func (u *LoginHandlerImpl) IsLogin(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "You are logged in"})
}
