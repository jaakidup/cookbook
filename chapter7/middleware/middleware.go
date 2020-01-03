package middleware

import "net/http"

import "log"

import "time"

// Middleware is what all middleware function will return
type Middleware func(http.HandlerFunc) http.HandlerFunc

// ApplyMiddleware will apply all middleware, the last
// argument will be the
// outer wrap for context passing purposes
func ApplyMiddleware(h http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
	applied := h

	// wrap around the middlewares
	for _, m := range middleware {
		applied = m(applied)
	}
	return applied
}

// Logger logs requests, this will use an id passed in via SetID()
func Logger(l *log.Logger) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			l.Printf("started request to %s with id %s ", r.URL, GetID(r.Context()))
			next(w, r)
			l.Printf("completed request to %s with id %s in %s", r.URL, GetID(r.Context()), time.Since(start))
		}
	}
}
