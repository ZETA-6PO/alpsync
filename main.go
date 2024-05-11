package main

import (
	"alpsync-api/api"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	fmt.Println("[alpsync-api] Server running...")

	// Definir les routes
	router := mux.NewRouter()
	api.InitRoutes(router)
	// Servir les fichier de facon statique
	http.Handle("/", http.FileServer(http.Dir("./static/")))

	certFile := "./ssl/server.crt"
	keyFile := "./ssl/server.key"

	// Demarrer le serveur HTTP avec HTTPS
	log.Fatal(http.ListenAndServeTLS(":8080", certFile, keyFile, router))
	// Demarrer le serveur avec HTTP
	//log.Fatal(http.ListenAndServe(":8080", router))

}
