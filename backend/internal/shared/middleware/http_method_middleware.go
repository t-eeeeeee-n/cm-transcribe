package middleware

import (
	"net/http"
)

// HttpMethodMiddleware リクエストメソッドをチェックするミドルウェア
func HttpMethodMiddleware(next http.Handler, expectedMethod string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != expectedMethod {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}
