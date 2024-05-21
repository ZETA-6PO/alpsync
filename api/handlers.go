package api

import (
	"alpsync-api/db"
	"alpsync-api/utils"
	"encoding/json"
	"fmt"
	"net/http"
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
		uploadErr(w, err.Error())
		return
	}

	// Get the file from the content-type multipart request
	file, handler, err := r.FormFile("file")
	if err != nil {
		uploadErr(w, err.Error())
		return
	}
	defer file.Close()

	// create a db entry for that file
	hexId, err := db.AddFileEntry(handler.Filename, handler.Header.Get("expiresAt"))
	if err != nil {
		fmt.Println("Error db : ", err.Error())
		uploadErr(w, err.Error())
		return
	}

	// create an file on disk
	file_path := "./data/" + hexId
	err = utils.CreateFile(file, file_path)
	if err != nil {
		uploadErr(w, err.Error())
		return
	}

	var footprint float64 = (7 / 365000) * (float64(handler.Size))

	uploadOk(w, hexId, handler.Filename, footprint)

}

func downloadPageHandler(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path

	parts := strings.Split(urlPath, "/")

	if len(parts) != 3 || parts[1] != "f" {
		downloadErr(w, "bad url format")
		fmt.Println("")
		fmt.Println("error")
		return
	}

	code := parts[2]

	filename, err := db.GetFileEntry(code)

	if err != nil {
		downloadErr(w, err.Error())
		return
	}

	fmt.Printf("filename is %s\n", filename)

	downloadOk(w, filename, code)

}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path

	parts := strings.Split(urlPath, "/")

	if len(parts) != 4 || parts[2] != "dwl" {
		fmt.Println("error")
		downloadErr(w, "bad url format")
		return
	}

	code := parts[3]

	filename, err := db.GetFileEntry(code)

	if err != nil {
		fmt.Println("error")
		downloadErr(w, err.Error())
		return
	}

	fileStat, file, err := utils.ReadFile("./data/" + code)

	defer file.Close()

	if err != nil {
		downloadErr(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	http.ServeContent(w, r, filename, fileStat.ModTime(), file)

}
