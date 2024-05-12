package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Service is available."}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// ParseMultipartForm analyse une requête avec un corps multipart et stocke les résultats dans r.MultipartForm.
	err := r.ParseMultipartForm(50 << 20) // 10 Mo de taille maximale de fichier
	if err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusBadRequest)
		return
	}

	// Récupère le fichier du champ de formulaire "file"
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Crée un fichier local pour stocker le fichier téléchargé
	f, err := os.OpenFile("./data/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Copie le contenu du fichier téléchargé dans le fichier local
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Failed to copy file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
}
