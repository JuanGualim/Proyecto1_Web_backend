package routes

import (
	"net/http"

	"series-api/handlers"
)

func SetupRoutes() {

	http.HandleFunc("/series", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetSeries(w, r)
		} else if r.Method == http.MethodPost {
			handlers.CreateSeries(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/series/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetSeriesByID(w, r)
		case http.MethodPut:
			handlers.UpdateSeries(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
