package notary

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/NTRYPlatform/ntry-backend/config"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

func Index(handler *Handler) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handler.data = "Welcome to Notary Platform"
			handler.ServeHTTP(w, r)
		})
	}
}

func NotImplemented(handler *Handler) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handler.data = "Function not implemented yet!"
			handler.ServeHTTP(w, r)
		})
	}
}

func Logging(handler *Handler) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handler.logger.Info(fmt.Sprintf("[ntry] new request %s: %s %s", r.UserAgent(), r.Method, r.URL))
			h.ServeHTTP(w, r)
		})
	}
}

func Authorization(handler *Handler) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		})
	}
}

func CreateUser(handler *Handler, email *emailConf) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			u := &User{}

			if err := decode(r, u); err != nil {
				// Set error data and jump to the last handler
				// implemented by *Handler
				handler.status = http.StatusInternalServerError
				handler.data = err
				handler.ServeHTTP(w, r)
			}

			// check for pre-exsitance
			if ok, err := handler.db.UserExistsByUniqueField(u); err != nil {

				handler.status = http.StatusInternalServerError
				handler.data = err
				handler.ServeHTTP(w, r)

			} else if ok {

				handler.logger.Error(
					fmt.Sprintf("[handler ] User with either of these values already exists!user: %v, err: %v", u, err))
				handler.status = http.StatusInternalServerError
				handler.data = "User already Exist"
				handler.ServeHTTP(w, r)
			}

			u.RegTime = time.Now().UTC()

			if err := handler.db.Insert(u); err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] User insertion to db error!user: %v, err: %v", u, err))
				handler.status = http.StatusInternalServerError
				handler.data = err
				handler.ServeHTTP(w, r)
			}

			handler.logger.Info(fmt.Sprint("[handler ] User successfully saved to db!", u.String()))

			msg := verificationAccountMessage(email.from, u.EmailAddress, u.UID)

			if err := email.sendEmail(u.EmailAddress, msg); err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Failed to send verification email! user: %v, err: %v", u, err))
				handler.status = http.StatusInternalServerError
				handler.data = err
				handler.ServeHTTP(w, r)
			}

			// Folow the normal flow
			handler.status = http.StatusOK
			handler.data = u
			w.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(w, r)
		})
	}
}

func UpdateUserInfo(handler *Handler) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u := &User{}

			if err := decode(r, u); err != nil {
				// Set error data and jump to the last handler
				// implemented by *Handler
				handler.status = http.StatusInternalServerError
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				handler.data = err
				handler.ServeHTTP(w, r)
			}

			if err := handler.db.UpdateUser(u); err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Failed to update user record!user: %v, err: %v", u, err))
				handler.status = http.StatusInternalServerError
				handler.data = err
				handler.ServeHTTP(w, r)
			}

			// Folow the normal flow
			handler.status = http.StatusCreated
			handler.data = true
			w.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(w, r)
		})
	}
}

func LoginHandler(handler *Handler, conf *config.Config) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u := &LoginUser{}
			user := &User{}

			var err error

			if err = decode(r, u); err != nil {
				// Set error data and jump to the last handler
				// implemented by *Handler
				handler.status = http.StatusForbidden
				handler.data = err
				handler.ServeHTTP(w, r)
			}

			if user, err = handler.db.LoginUserValidation(u); err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Failed to verify user! user: %v, err: %v", u, err))
				handler.status = http.StatusForbidden
				handler.data = err
				handler.ServeHTTP(w, r)
			}

			if user == nil {
				handler.logger.Info(
					fmt.Sprintf("[handler ] Ivalid credentials! user: %v", user))
				handler.status = http.StatusForbidden
				handler.data = "Invalid credentials"
				handler.ServeHTTP(w, r)
			}

			//TODO: trigger eth network check

			// Create a map to store user claims
			//TODO: from db
			claims := UserJWT{
				*user,
				jwt.StandardClaims{
					IssuedAt: time.Now().Unix(),
					// expires in an hour
					ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
					//TODO: change in production - should be configurable
					Issuer: conf.GetServerAddress(),
				},
			}

			// Create token with claims
			// TODO: might want to use signing method ECDSA with pvt key
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			pvtKey, err := conf.GetPvtKey()
			if err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Unable to get private key! user: %v, err: %v", user, err))
				handler.status = http.StatusInternalServerError
				handler.data = "Server Error" // TODO: may be some meaningful message
				handler.ServeHTTP(w, r)
			}

			tokenString, err := token.SignedString(pvtKey)

			if err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Error while signing the token! user: %v, err: %v", user, err))
				handler.status = http.StatusInternalServerError
				handler.data = "Error while signing the token"
				handler.ServeHTTP(w, r)
			}

			json, err := json.Marshal(tokenString)
			if err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Error in json marshaling of token string! user: %v, err: %v", user, err))
				handler.status = http.StatusInternalServerError
				handler.data = "Internal Error"
				handler.ServeHTTP(w, r)
			}

			// Follow the normal flow
			handler.status = http.StatusOK
			handler.data = json
			w.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(w, r)

		})
	}
}

/* --------------- Middlewares ---------------- */

func ValidateTokenMiddleware(handler *Handler, conf *config.Config) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			pubKey, err := conf.GetPubKey()
			if err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Unable to get public key! err: %v", err))
				handler.status = http.StatusInternalServerError
				handler.data = "Unexpected Error"
				handler.ServeHTTP(w, r)
			}

			token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
				func(token *jwt.Token) (interface{}, error) {
					return pubKey, nil
				})

			if err == nil {
				if token.Valid {
					h.ServeHTTP(w, r)
				} else {
					handler.logger.Info(
						fmt.Sprint("[handler ] Invalid Token!"))
					handler.status = http.StatusForbidden
					handler.data = "Not allowed"
					handler.ServeHTTP(w, r)
				}
			} else {
				handler.logger.Error(
					fmt.Sprint("[handler ] Unauthorized access!"))
				handler.status = http.StatusForbidden
				handler.data = "Unauthorized Access"
				handler.ServeHTTP(w, r)
			}
		})
	}
}

// AuthMiddleware
func AuthMiddleware(handler *Handler, conf *config.Config) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			u := &User{}

			// Create a map to store user claims
			//TODO: from db
			claims := UserJWT{
				*u,
				jwt.StandardClaims{
					Id:       "someuserid",
					IssuedAt: time.Now().Unix(),
					// expires in an hour
					ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
					//TODO: change in production - should be configurable
					Issuer: conf.GetServerAddress(),
				},
			}

			// Create token with claims
			// TODO: might want to use signing method ECDSA with pvt key
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			pvtKey, err := conf.GetPvtKey()
			if err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Unable to get private key! user: %v, err: %v", u, err))
				handler.status = http.StatusInternalServerError
				handler.data = "Server Error" // may be some meaningful message
				handler.ServeHTTP(w, r)
			}

			// Sign the token with our secret
			tokenString, _ := token.SignedString(pvtKey)

			handler.logger.Info(
				fmt.Sprintf("[handler ] Token string signed successfullly! user: %v, TokenString: %v", u, tokenString))

			h.ServeHTTP(w, r)
		})
	}
}
