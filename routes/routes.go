package routes

import (
	"net/http"

	"series-api/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/series", handlers.GetSeries)
}
