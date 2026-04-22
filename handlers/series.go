package handlers

import (
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
