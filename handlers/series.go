package handlers

import (
	"encoding/json"
	"net/http"

	"series-api/database"
	"series-api/models"
	"series-api/utils"
)

func GetSeries(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(w)

	rows, err := database.DB.Query("SELECT id, name, current_episode, total_episodes, image_url FROM series")
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Error fetching series")
		return
	}

	if r.Method == http.MethodOptions {
		return
	}
	defer rows.Close()

	var seriesList []models.Series

	for rows.Next() {
		var s models.Series
		err := rows.Scan(&s.ID, &s.Name, &s.CurrentEpisode, &s.TotalEpisodes, &s.ImageURL)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, "Error reading data")
			return
		}
		seriesList = append(seriesList, s)
	}

	utils.JSON(w, http.StatusOK, seriesList)
}

func CreateSeries(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(w)

	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodPost {
		utils.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var s models.Series

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	// 🔍 Validaciones
	if s.Name == "" {
		utils.Error(w, http.StatusBadRequest, "Name is required")
		return
	}

	if s.CurrentEpisode < 1 {
		utils.Error(w, http.StatusBadRequest, "Current episode must be >= 1")
		return
	}

	if s.TotalEpisodes < 1 {
		utils.Error(w, http.StatusBadRequest, "Total episodes must be >= 1")
		return
	}

	if s.CurrentEpisode > s.TotalEpisodes {
		utils.Error(w, http.StatusBadRequest, "Current episode cannot exceed total episodes")
		return
	}

	// 💾 Insertar en DB
	query := `
	INSERT INTO series (name, current_episode, total_episodes, image_url)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`

	err = database.DB.QueryRow(
		query,
		s.Name,
		s.CurrentEpisode,
		s.TotalEpisodes,
		s.ImageURL,
	).Scan(&s.ID)

	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Error inserting series")
		return
	}

	utils.JSON(w, http.StatusCreated, s)
}
