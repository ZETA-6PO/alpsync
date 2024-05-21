package api

import (
	"alpsync-api/db"
	"alpsync-api/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Handle get status request
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

// Handle file upload request
func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(50 << 20)
	if err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusBadRequest)
		return
	}

	// Get the file from the content-type multipart request
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// create a db entry for that file
	hexId, err := db.AddFileEntry(handler.Filename, handler.Header.Get("expiresAt"))
	if err != nil {
		fmt.Println("Error db : ", err.Error())
		http.Error(w, "Failed to create entry in db", http.StatusInternalServerError)
	}

	// create an file on disk
	file_path := "./data/" + hexId
	err = utils.CreateFile(file, file_path)
	if err != nil {
		http.Error(w, "Failed to write the file on disk.", http.StatusInternalServerError)
	}

	baseURL := "https://alpsync.pro/uploadOk"

	// Création d'un objet url.Values pour stocker les paramètres de requête
	params := url.Values{}
	params.Set("fileid", "https://alpsync.pro/f-"+hexId)
	params.Set("filename", handler.Filename)
	params.Set("footprint", "2.5g")
	http.Redirect(w, r, baseURL+"?"+params.Encode(), http.StatusSeeOther)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	// Recuperer le code de l'URL
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 || parts[1] != "f-" {
		http.NotFound(w, r)
		return
	}
	code := parts[2]
	//Verifier si le fichier existe dans la bdd
	filename, err := db.GetFileEntry(code)

	if err != nil {
		http.Error(w, "Unknown file id.", http.StatusBadRequest)
		return
	}

	fileStat, file, err := utils.ReadFile("./data/" + code)

	if err != nil {
		http.Error(w, "File cannot be opened on the server.", http.StatusInternalServerError)
		return
	}

	// Set le header Content-Type pour indiquer le type de fichier
	w.Header().Set("Content-Type", "application/octet-stream")

	// Set le header Content-Disposition pour indiquer le nom du fichier lors du téléchargement
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	// Copier le contenu du fichier dans la réponse HTTP
	http.ServeContent(w, r, filename, fileStat.ModTime(), file)
}
