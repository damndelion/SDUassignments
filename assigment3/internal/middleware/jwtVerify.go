package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

func JwtVerify(secretKey string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenHeader := r.Header.Get("Authorization")
			if tokenHeader == "" {
				http.Error(w, "No token provided", http.StatusUnauthorized)
				return
			}

			token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(secretKey), nil
			})

			var claims jwt.MapClaims
			var ok bool
			if claims, ok = token.Claims.(jwt.MapClaims); ok && token.Valid {
				if exp, ok := claims["exp"].(float64); ok {
					expirationTime := time.Unix(int64(exp), 0)
					if time.Now().After(expirationTime) {
						http.Error(w, err.Error(), http.StatusUnauthorized)
						return
					}
				} else {
					http.Error(w, err.Error(), http.StatusUnauthorized)
					return
				}
			} else {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)

		})
	}
}
