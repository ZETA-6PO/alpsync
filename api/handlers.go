package api

import (
	"alpsync-api/db"
	"alpsync-api/utils"
	"encoding/json"
	"fmt"
	"net/http"
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
	file_path := "./data/" + hexId + "_" + handler.Filename
	err = utils.CreateFile(file, file_path)
	if err != nil {
		http.Error(w, "Failed to write the file on disk.", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename) // must be remove
}
