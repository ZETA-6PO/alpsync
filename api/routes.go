package api

import (
	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {

	router.HandleFunc("/uploadOk", uploadOk)
	router.PathPrefix("/f").Subrouter().HandleFunc("/", downloadHandler)
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/status", statusHandler).Methods("GET")
	apiRouter.HandleFunc("/upload", uploadFileHandler).Methods("POST")

}
