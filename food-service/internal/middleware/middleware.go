package middleware

import (
	"context"
	"net/http"
)

func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAuth := struct{ key string }{key: r.Header.Get("Authorization")}
		ctx := r.Context()
		if userAuth.key != "" {
			ctx = context.WithValue(ctx, userAuth, userAuth)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
