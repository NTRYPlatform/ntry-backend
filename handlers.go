package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/ntryapp/auth/config"
)

func Index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Welcome to Notary!")
}

func NotImplemented(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Function not implemented yet!")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	//TODO: check the limit
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RegisterUser(user)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	var user User
	//TODO: check the limit
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := CompleteUserInfo(&user)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func VerifySecondaryAddress(w http.ResponseWriter, r *http.Request) {
	var user VerifyUserSignature
	//TODO: check the limit
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := ValidateSecondaryAddress(&user)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var user LoginUser

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error in request")
		return
	}

	u := LoginUserValidation(&user)

	if u == nil {
		w.WriteHeader(http.StatusForbidden)
		log.Println("Error logging in! Invalid credentials.")
		return
	}

	//TODO: trigger eth network check

	// Create a map to store user claims
	//TODO: from db
	claims := UserJWT{
		*u,
		jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
			// expires in an hour
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			//TODO: change in production - should be configurable
			Issuer: config.GetServerAddress(),
		},
	}

	// Create token with claims
	// TODO: might want to use signing method ECDSA with pvt key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	fmt.Fprintln(w, "Error extracting the key")
	// }

	tokenString, err := token.SignedString(config.GetPvtKey())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(w, "Error while signing the token")
	}

	json, err := json.Marshal(tokenString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

}

/* --------------- Middlewares ---------------- */

func ValidateTokenMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				return config.GetPubKey(), nil
			})

		if err == nil {
			if token.Valid {
				h.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Token is not valid")
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized access to this resource")
		}
	})
}

// AuthMiddleware
func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("This ought to be fun!!!")
		u := User{}

		// Create a map to store user claims
		//TODO: from db
		claims := UserJWT{
			u,
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
		tokenString, _ := token.SignedString(config.GetPvtKey())

		log.Println(tokenString)
		h.ServeHTTP(w, r)
	})
}
