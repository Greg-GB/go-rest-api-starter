package middleware

import (
	"context"
	"math/rand"
	"net/http"
	"strconv"
)

func AddRequestID(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// The request id should be unique so consider using a uuid.
		// github.com/docker/distribution/uuid
		ctx := context.WithValue(r.Context(), "request-id", strconv.Itoa(rand.Int()))
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
		return
	})
}

func Wrapper(h http.Handler, middlewares ...middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}
