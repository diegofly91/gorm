package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"gorm/src/common/handlers"
	"gorm/src/modules/auth/models"
	authClaims "gorm/src/modules/auth/models"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte("your_secret_key")

type contextKey string

const userContextKey = contextKey("user")

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			handlers.SendError(rw, http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.ParseWithClaims(tokenString, &authClaims.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return JwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(rw, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*authClaims.CustomClaims)
		if !ok {
			handlers.SendError(rw, http.StatusUnauthorized)
			return
		}
		fmt.Println("Username", claims.Username)
		fmt.Println("Email", claims.Email)

		ctx := context.WithValue(r.Context(), "user", models.UserPayload{
			Username: claims.Username,
			Email:    claims.Email,
		})

		r = r.WithContext(ctx)

		next(rw, r)
	}
}
