package middleware

import (
	"net/http"

	"github.com/AndresMaldoando24/go_htmx_base/handlers"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := handlers.Store.Get(r, "session-name")
		user := session.Values["user_id"]

		if err != nil || user == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)

	})
}
