package main

import (
	"fmt"
	"net/http"

	"series-api/database"
	"series-api/routes"
)

func main() {

	// DB
	database.Connect()
	database.InitTables()
	routes.SetupRoutes()

	http.Handle("/swagger.yaml", http.FileServer(http.Dir(".")))

	http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("swagger-ui"))))

	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	// Ruta raíz
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API running 🚀")
	})

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", corsMiddleware(http.DefaultServeMux))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
