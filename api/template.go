package api

import (
	"fmt"
	"html/template"
	"net/http"
)

// Structure pour stocker les données
type DataUploadOk struct {
	FileName  string
	FileId    string
	FootPrint float64
}

type DataUploadErr struct {
	Message string
}

type DataDownloadErr struct {
	Message string
}

type DataDownloadOk struct {
	FileName string
}

func uploadOk(w http.ResponseWriter, r *http.Request) {
	// Retrieve values from the URL query string
	fileID := r.URL.Query().Get("fileid")
	filename := r.URL.Query().Get("filename")
	footprint := r.URL.Query().Get("footprint")

	// Convert footprint to float64
	var footPrintFloat float64
	i, err := fmt.Sscanf(footprint, "%f", &footPrintFloat)
	if i != 1 || err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a Data object with retrieved values
	data := DataUploadOk{
		FileId:    fileID,
		FileName:  filename,
		FootPrint: footPrintFloat,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Load the template file
	tmpl, err := template.ParseFiles("./template/uploadOk.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template using the data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func uploadErr(w http.ResponseWriter, r *http.Request) {
	// Retrieve values from the URL query string
	message := r.URL.Query().Get("message")

	data := DataUploadErr{
		Message: message,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Load the template file
	tmpl, err := template.ParseFiles("./template/uploadErr.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template using the data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func downloadOk(w http.ResponseWriter, r *http.Request) {
	// Retrieve values from the URL query string
	filename := r.URL.Query().Get("filename")

	data := DataDownloadOk{
		FileName: filename,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Load the template file
	tmpl, err := template.ParseFiles("./template/downloadOk.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template using the data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func downloadErr(w http.ResponseWriter, message string) {

	data := DataDownloadErr{
		Message: message,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Load the template file
	tmpl, err := template.ParseFiles("./template/downloadErr.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template using the data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
