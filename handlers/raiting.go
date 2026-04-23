package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"series-api/database"
	"series-api/utils"
)

type RatingInput struct {
	Rating int `json:"rating"`
}

func AddRating(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(w)

	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodPost {
		utils.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Extraer ID
	path := strings.TrimPrefix(r.URL.Path, "/series/")
	parts := strings.Split(path, "/")

	if len(parts) < 2 {
		utils.Error(w, http.StatusBadRequest, "Invalid URL")
		return
	}

	id, err := strconv.Atoi(parts[0])
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var input RatingInput

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	// 🔍 Validación
	if input.Rating < 1 || input.Rating > 5 {
		utils.Error(w, http.StatusBadRequest, "Rating must be between 1 and 5")
		return
	}

	// Insertar rating
	query := `
	INSERT INTO ratings (series_id, rating)
	VALUES ($1, $2)
	`

	_, err = database.DB.Exec(query, id, input.Rating)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Error inserting rating")
		return
	}

	utils.JSON(w, http.StatusCreated, map[string]string{
		"message": "Rating added",
	})

}
