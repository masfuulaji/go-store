package utils

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func ExtractUserIDFromJWT(r *http.Request) (int, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return 0, fmt.Errorf("missing token cookie")
	}
	tokenString := cookie.Value

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})
	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	return claims.Id, nil
}
