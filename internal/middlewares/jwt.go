package middlewares

import (
	"context"
	"database/sql"
	"github.com/iyorozuya/real-world-app/internal/utils"
	"net/http"
)

func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		userId, err := utils.ParseAuthorizationHeader(tokenHeader)
		if err != nil {
			utils.SendError(w, http.StatusForbidden, "access forbidden")
			return
		}
		ctx := context.WithValue(r.Context(), "userId", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func OptionalJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		userId, err := utils.ParseAuthorizationHeader(tokenHeader)
		var ctx context.Context
		if err != nil {
			ctx = context.WithValue(r.Context(), "userId", sql.NullString{})
		} else {
			ctx = context.WithValue(r.Context(), "userId", sql.NullString{
				String: userId,
				Valid:  true,
			})
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
