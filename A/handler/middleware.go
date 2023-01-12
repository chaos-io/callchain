package handler

import (
	"net/http"

	"github.com/chaos-io/go-trace"
)

func MiddlewareTrace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		trace.Trace()

		next.ServeHTTP(w, r)
	})
}
