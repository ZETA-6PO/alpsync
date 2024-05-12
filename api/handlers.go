package api

import (
	"alpsync-api/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/kamva/mgm/v3"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
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
	err := r.ParseMultipartForm(50 << 20)
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

	//creer une entree dans la base de donne
	dbentry := models.NewASFile(handler.Filename, time.Now().Format(time.UnixDate))
	coll := mgm.CollectionByName("files")
	err = coll.Create(dbentry)
	if err != nil {
		fmt.Println("Error db : ", err.Error())
		http.Error(w, "Failed to create entry in db", http.StatusInternalServerError)
	}
	file_path := "./data/" + dbentry.ID.Hex() + "_" + handler.Filename

	// Crée un fichier local pour stocker le fichier téléchargé
	f, err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE, 0666)
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
