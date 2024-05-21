package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitRoutes(router *mux.Router) {
	// Servir les fichier de facon statique
	router.PathPrefix("/f/").HandlerFunc(downloadHandler)
	fs := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/").Handler(fs)
	router.HandleFunc("/uploadOk", uploadOk)
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/status", statusHandler).Methods("GET")
	apiRouter.HandleFunc("/upload", uploadFileHandler).Methods("POST")

}
