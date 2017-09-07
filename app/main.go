package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ntryapp/auth"
	"github.com/ntryapp/auth/config"
)

func main() {

	addr := config.GetServerAddress()

	// main router
	router := mux.NewRouter().StrictSlash(false)
	// api sub-router
	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/", auth.Index)
	router.HandleFunc("/sign-up", auth.CreateUser)
	//TODO: technically should be restricted
	router.HandleFunc("/update-info", auth.UpdateUserInfo)
	router.HandleFunc("/verify-address", auth.VerifySecondaryAddress)
	router.HandleFunc("/sign-in", auth.LoginHandler)
	router.HandleFunc("/send-invitation", auth.NotImplemented)

	http.Handle("/", router)
	http.Handle("/api", auth.AuthMiddleware(apiRouter))

	//TODO:generate cert and serve on TLS
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Panicln("Whaaaaa?", err)
	}
}
