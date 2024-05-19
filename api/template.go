package api

import (
	"fmt"
	"html/template"
	"net/http"
)

// Structure pour stocker les données
type Data struct {
	FileName  string
	FileId    string
	FootPrint float64
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
	data := Data{
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
