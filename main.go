package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	LoadConfig()
	addr := GetServerAddress()

	// main router
	router := mux.NewRouter().StrictSlash(false)
	// api sub-router
	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/", Index)
	router.HandleFunc("/sign-up", CreateUser)
	router.HandleFunc("/sign-in", LoginHandler)
	router.HandleFunc("/send-invitation", NotImplemented)

	http.Handle("/", router)
	http.Handle("/api", AuthMiddleware(apiRouter))

	//TODO:generate cert and serve on TLS
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Panicln("Whaaaaa?", err)
	}
}
