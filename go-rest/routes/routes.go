package routes

import (
	"net/http"

	"github.com/manish-npx/go-lang/go-rest/controllers"
)

func RegisterRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Write([]byte("Manish ====> Yadav"))
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controllers.GetUsers(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controllers.GetUserByID(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
