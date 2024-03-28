package middleware

import (
	"context"
	"log"
	"net/http"
	"regexp"

	"github.com/krishnapramodaradhi/expense-tracker-api/internal/util"
)

var tokenRegex = regexp.MustCompile(`Bearer [a-zA-z0-9.]+`)

func TokenValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			util.WriteJSON(w, http.StatusUnauthorized, map[string]string{"error": "Token not found in the request"})
			return
		}
		if !tokenRegex.Match([]byte(tokenString)) {
			util.WriteJSON(w, http.StatusUnauthorized, map[string]string{"error": "Not a Bearer token"})
			return
		}
		userId, err := util.ValidateToken(tokenString)
		if err != nil {
			log.Println(err)
			return
		}
		ctx := context.WithValue(r.Context(), "userId", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
