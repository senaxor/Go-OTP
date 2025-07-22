package middlewares

import (
	"log"
	"net/http"
	"os"
)

func AdminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			log.Println("AdminAuth: missing or invalid basic auth header")
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		adminUser := os.Getenv("ADMIN_USERNAME")
        adminPass := os.Getenv("ADMIN_PASSWORD")

        if username != adminUser || password != adminPass {
            log.Printf("AdminAuth: unauthorized login attempt with username: %s\n", username)
            w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        log.Printf("AdminAuth: successful login for user: %s\n", username)
        next.ServeHTTP(w, r)
	})
}
