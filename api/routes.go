package api

import (
	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/test", helloWorldHandler).Methods("GET")

}
