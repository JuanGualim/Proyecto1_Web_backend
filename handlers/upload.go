package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"series-api/utils"
)

func UploadImage(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(w)

	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodPost {
		utils.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Parsear form
	err := r.ParseMultipartForm(2 << 20) // 2MB max
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "File too large")
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Error reading file")
		return
	}
	defer file.Close()

	// Generar nombre único
	ext := filepath.Ext(handler.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	// Crear archivo destino
	dst, err := os.Create("uploads/" + filename)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Error saving file")
		return
	}
	defer dst.Close()

	// Copiar contenido
	_, err = io.Copy(dst, file)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Error writing file")
		return
	}

	// URL pública
	imageURL := fmt.Sprintf("http://localhost:8080/uploads/%s", filename)

	utils.JSON(w, http.StatusCreated, map[string]string{
		"url": imageURL,
	})
}
