package main

import (
	"log"
	"net/http"
	"ntry-user-mgmt/config"

	"github.com/gorilla/mux"
)

func init() {
	// Load configuration
	config.LoadConfig()
}

func main() {

	addr := config.GetServerAddress()

	// main router
	router := mux.NewRouter().StrictSlash(false)
	// api sub-router
	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/", Index)
	router.HandleFunc("/sign-up", CreateUser)
	//TODO: technically should be restricted
	router.HandleFunc("/update-info", UpdateUserInfo)
	router.HandleFunc("/verify-address", VerifySecondaryAddress)
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
