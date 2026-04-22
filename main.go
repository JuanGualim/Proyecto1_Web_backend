package main

import (
	"fmt"
	"net/http"
	"series-api/routes"

	"series-api/database"
)

func main() {
	// Conectar DB
	database.Connect()
	database.InitTables()
	routes.SetupRoutes()

	// Ruta de prueba
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(&w)
		fmt.Fprintf(w, "API running 🚀")
	})

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
