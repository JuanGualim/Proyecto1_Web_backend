package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

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

func GetSeriesByID(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(w)

	if r.Method == http.MethodOptions {
		return
	}

	// Obtener ID de la URL
	idStr := strings.TrimPrefix(r.URL.Path, "/series/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var s models.Series

	query := `
	SELECT id, name, current_episode, total_episodes, image_url
	FROM series
	WHERE id = $1
	`

	err = database.DB.QueryRow(query, id).Scan(
		&s.ID,
		&s.Name,
		&s.CurrentEpisode,
		&s.TotalEpisodes,
		&s.ImageURL,
	)

	if err != nil {
		utils.Error(w, http.StatusNotFound, "Series not found")
		return
	}

	utils.JSON(w, http.StatusOK, s)
}

func UpdateSeries(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(w)

	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodPut {
		utils.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Obtener ID
	idStr := strings.TrimPrefix(r.URL.Path, "/series/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var s models.Series

	err = json.NewDecoder(r.Body).Decode(&s)
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

	// 🛠 Update
	query := `
	UPDATE series
	SET name=$1, current_episode=$2, total_episodes=$3, image_url=$4
	WHERE id=$5
	`

	result, err := database.DB.Exec(
		query,
		s.Name,
		s.CurrentEpisode,
		s.TotalEpisodes,
		s.ImageURL,
		id,
	)

	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Error updating series")
		return
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		utils.Error(w, http.StatusNotFound, "Series not found")
		return
	}

	s.ID = id
	utils.JSON(w, http.StatusOK, s)
}

func DeleteSeries(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(w)

	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodDelete {
		utils.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Obtener ID
	idStr := strings.TrimPrefix(r.URL.Path, "/series/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	// 🗑 Eliminar
	query := `DELETE FROM series WHERE id = $1`

	result, err := database.DB.Exec(query, id)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Error deleting series")
		return
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		utils.Error(w, http.StatusNotFound, "Series not found")
		return
	}

	// ✔ 204 No Content (correcto en REST)
	w.WriteHeader(http.StatusNoContent)
}
