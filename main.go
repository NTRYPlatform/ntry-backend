package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var mySigningKey = []byte("secret")

func Index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Welcome to Notary!")
}

func authMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("This ought to be fun!!!")

		// Create a map to store user claims
		//TODO: from db
		claims := UserClaims{
			"0xfakeethereumadress",
			"password",
			jwt.StandardClaims{
				Id:       "someuserid",
				IssuedAt: time.Now().Unix(),
				// expires in an hour
				ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
				//TODO: change in production - should be configurable
				Issuer: "localhost:9090",
			},
		}

		// Create token with claims
		// TODO: might want to use signing method ECDSA with pvt key
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Sign the token with our secret
		tokenString, _ := token.SignedString(mySigningKey)

		log.Println(tokenString)
		h.ServeHTTP(w, r)
	})
}

func main() {

	addr := GetServerAddress()

	fmt.Println(addr)

	//api router
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	http.Handle("/", authMiddleware(router))

	//TODO:generate cert and serve on TLS
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Panicln("whaaaaa?", err)
	}
}
