package routes

import (
	"net/http"
	"strings"

	"series-api/handlers"
)

func SetupRoutes() {

	// 🔹 /series (GET + POST)
	http.HandleFunc("/series", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetSeries(w, r)
		case http.MethodPost:
			handlers.CreateSeries(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// 🔹 /series/... (ID + rating)
	http.HandleFunc("/series/", func(w http.ResponseWriter, r *http.Request) {

		// ⭐ manejar rating
		if strings.HasSuffix(r.URL.Path, "/rating") {

			if r.Method == http.MethodPost {
				handlers.AddRating(w, r)
				return
			}

			if r.Method == http.MethodGet {
				handlers.GetRating(w, r)
				return
			}
		}

		// 🔹 resto (ID)
		switch r.Method {
		case http.MethodGet:
			handlers.GetSeriesByID(w, r)
		case http.MethodPut:
			handlers.UpdateSeries(w, r)
		case http.MethodDelete:
			handlers.DeleteSeries(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
