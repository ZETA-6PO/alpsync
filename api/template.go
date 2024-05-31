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
	FootPrint string
	Duration  string
}

type DataUploadErr struct {
	Message string
}

type DataDownloadErr struct {
	Message string
}

type DataDownloadOk struct {
	FileName string
	FileId   string
}

func uploadOk(w http.ResponseWriter, fileID string, filename string, footprint float64, duration int) {

	footprintStr := fmt.Sprintf("%.2f", footprint)

	// Create a Data object with retrieved values
	data := DataUploadOk{
		FileId:    fileID,
		FileName:  filename,
		FootPrint: footprintStr,
		Duration:  fmt.Sprintf("%d", duration),
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

func uploadErr(w http.ResponseWriter, message string) {
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

func downloadOk(w http.ResponseWriter, filename string, fileid string) {

	data := DataDownloadOk{
		FileName: filename,
		FileId:   fileid,
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
