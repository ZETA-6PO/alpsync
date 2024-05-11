package api

import (
	"encoding/json"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hi !"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {

}
