package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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

	// Parsear form (máximo permitido por request)
	err := r.ParseMultipartForm(2 << 20) // 2MB
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

	if handler.Size > 1<<20 {
		utils.Error(w, http.StatusBadRequest, "File too large (max 1MB)")
		return
	}

	// 🔥 VALIDACIÓN DE TIPO (opcional pero recomendado)
	contentType := handler.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		utils.Error(w, http.StatusBadRequest, "Only image files are allowed")
		return
	}

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
