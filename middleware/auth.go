package middleware

import (
	"context"
	"fmt"
	"movie_api/utils"
	"net/http"
	"strings"
)

type contextKey string

const (
	UserIDKey contextKey = "userID"
	RoleKey   contextKey = "role"
)

func MiddlewareJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) == 2 && tokenParts[0] == "Bearer" {
				tokenString := tokenParts[1]
				claims, err := utils.ValidateJWT(tokenString)
				if err == nil {
					ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
					ctx = context.WithValue(ctx, RoleKey, claims.Role)
					r = r.WithContext(ctx)
					fmt.Println("✅ JWT valid. Claims:", claims)
				} else {
					fmt.Println("⚠️ Invalid JWT:", err)
				}
			}
		}
		// lanjut ke handler, entah token valid atau tidak
		next.ServeHTTP(w, r)
	})
}
