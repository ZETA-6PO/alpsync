package main

import (
	"alpsync-api/api"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

func main() {
	// Se connecter a la BDD
	log.Println("[alpsync-api] Connecting to MongoDB ", os.Getenv("ALPSYNC_MONGO_URL"))
	err := mgm.SetDefaultConfig(nil, "production", options.Client().ApplyURI(os.Getenv("ALPSYNC_MONGO_URL")))
	if err != nil {
		log.Fatal("[alpsync-api] Cannot connect to the database. ", err.Error())
	}

	fmt.Println("[alpsync-api] Server running...")

	// Definir les routes
	router := mux.NewRouter()
	api.InitRoutes(router)

	// Servir les fichier de facon statique
	fs := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/").Handler(fs)

	certFile := "./ssl/server.crt"
	keyFile := "./ssl/server.key"

	// Demarrer le serveur HTTP avec HTTPS
	log.Fatal(http.ListenAndServeTLS(":443", certFile, keyFile, router))
	// Demarrer le serveur avec HTTP
	//log.Fatal(http.ListenAndServe(":8080", router))

}
