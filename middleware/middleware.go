package middleware

import (
	"fmt"
	"net/http"

	"github.com/AndresMaldoando24/go_htmx_base/handlers"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := handlers.Store.Get(r, "session-name")
		token := session.Values["TOKEN"]

		fmt.Println(token)

		if err != nil || token == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)

	})
}
