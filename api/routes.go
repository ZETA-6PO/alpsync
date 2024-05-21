package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitRoutes(router *mux.Router) {
	// Servir les fichier de facon statique
	router.PathPrefix("/f/").HandlerFunc(downloadPageHandler)
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/status", statusHandler).Methods("GET")
	apiRouter.HandleFunc("/upload", uploadFileHandler).Methods("POST")
	apiRouter.HandleFunc("/dwl/", downloadHandler).Methods("GET")
	fs := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/").Handler(fs)

}
