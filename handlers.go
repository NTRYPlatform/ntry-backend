package notary

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/NTRYPlatform/ntry-backend/config"
	"github.com/NTRYPlatform/ntry-backend/eth"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type Token struct {
	Token string `json:"token"`
}

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

//TODO
func Authorization(handler *Handler) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		})
	}
}

func CreateUser(handler *Handler, email *emailConf, conf *config.Config) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u := &User{}
			if err := decode(r, u); err != nil {
				// Set error data and jump to the last handler
				// implemented by *Handler
				handler.status = http.StatusBadRequest
				handler.data = err
				handler.ServeHTTP(w, r)
				return
			}

			// check if user is valid
			if err := u.OK(); err != nil {
				handler.status = http.StatusBadRequest
				handler.data = err
				handler.ServeHTTP(w, r)
				return
			}
			// Add 0x to user id
			if u.UID[:2] != "0x" {
				u.UID = "0x" + u.UID
			}

			// check for pre-existence
			if exists, err := handler.db.UserExistsByUniqueField(u); err != nil {
				handler.status = http.StatusInternalServerError
				handler.data = err
				handler.ServeHTTP(w, r)
				return

			} else if exists {
				handler.logger.Error(
					fmt.Sprintf("[handler ] User with either of these values already exists!user: %v, err: %v", u, err))
				handler.status = http.StatusInternalServerError
				handler.data = "User already Exists"
				handler.ServeHTTP(w, r)
				return
			} else {
				cTime := time.Now().UTC()
				u.RegTime = &cTime
				u.AccountVerified = false
				fmt.Printf("Password: %s", u.Password)
				if u.Password, err = HashPassword(u.Password); err != nil {
					handler.logger.Error(
						fmt.Sprintf("[handler ] Couldn't hash password! user: %v, err: %v", u, err))
					handler.status = http.StatusInternalServerError
					handler.data = err
					handler.ServeHTTP(w, r)
					return
				}

				if err := handler.db.Insert(u, UserCollection); err != nil {
					handler.logger.Error(
						fmt.Sprintf("[handler ] User insertion to db error! user: %v, err: %v", u, err))
					handler.status = http.StatusInternalServerError
					handler.data = err
					handler.ServeHTTP(w, r)
					return
				}

				handler.logger.Info(fmt.Sprint("[handler ] User successfully saved to db!", u.String()))

				msg := verificationAccountMessage(email.from, u.EmailAddress, u.UID, conf.GetMapperContract())

				if err := email.sendEmail(u.EmailAddress, msg); err != nil {
					handler.logger.Error(
						fmt.Sprintf("[handler ] Failed to send verification email! user: %v, err: %v", u, err))
					handler.status = http.StatusInternalServerError
					handler.data = err
					handler.ServeHTTP(w, r)
					return
				}

				handler.status = http.StatusCreated
				handler.data = u.UID
			}

			h.ServeHTTP(w, r)

		})
	}
}

//TODO: shouldn't update everything/
func UpdateUserInfo(handler *Handler) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u := &User{}
			if err := decode(r, u); err != nil {
				// Set error data and jump to the last handler
				// implemented by *Handler
				handler.status = http.StatusBadRequest
				// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				handler.data = err
				handler.ServeHTTP(w, r)
				return
			}
			uid := context.Get(r, "uid")
			(*u).UID = uid.(string)
			if err := handler.db.UpdateUser(u); err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Failed to update user record!user: %v, err: %v", u, err))
				handler.status = http.StatusInternalServerError
				handler.data = err
				handler.ServeHTTP(w, r)
				return
			}
			// Follow the normal flow
			handler.status = http.StatusOK
			handler.data = true
			w.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(w, r)
			return

		})
	}
}

func ForgotPassword(handler *Handler, email *emailConf, conf *config.Config) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u := &LoginUser{}
			if err := decode(r, u); err != nil {
				// Set error data and jump to the last handler
				// implemented by *Handler
				handler.status = http.StatusBadRequest
				handler.data = err
				handler.ServeHTTP(w, r)
				return
			}

			// TODO: check if email is valid
			if len(u.EmailAddress) < 1 {
				handler.status = http.StatusInternalServerError
				handler.data = "Invalid email address!"
				handler.ServeHTTP(w, r)
				return
			}

			// generate temp password
			tmp := RandString(6)
			tmpHash, _ := HashPassword(tmp) //TODO: error check
			t := time.Now().Unix()
			if err := handler.db.InsertForgottenPassword(u.EmailAddress, tmpHash, t); err != nil {
				handler.status = http.StatusInternalServerError
				handler.data = "Can't provide account recovery!"
				handler.ServeHTTP(w, r)
				return

			}

			msg := changePasswordMessage(email.from, u.EmailAddress, time.Unix(t, 0).String(), tmp)

			if err := email.sendEmail(u.EmailAddress, msg); err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Failed to send verification email! user: %v, err: %v", u, err))
				handler.status = http.StatusInternalServerError
				handler.data = err
				handler.ServeHTTP(w, r)
				return
			}

			handler.status = http.StatusCreated // TODO: change to OK
			handler.data = "Check your email!"

			h.ServeHTTP(w, r)

		})
	}
}

//TODO: Need to add time check on subsequent tokens to make sure they're not older
func ChangePassword(handler *Handler) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u := &ChangePasswordUser{}
			if err := decode(r, u); err != nil {
				// Set error data and jump to the last handler
				// implemented by *Handler
				handler.status = http.StatusBadRequest
				handler.data = err
				handler.ServeHTTP(w, r)
				return
			}

			if err := u.OK(); err != nil {
				handler.status = http.StatusBadRequest
				handler.data = err
				handler.ServeHTTP(w, r)
				return
			}

			if err := handler.db.ChangeUserPassword(u); err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Failed to change password!user: %v, err: %v", u.EmailAddress, err))
				handler.status = http.StatusInternalServerError
				handler.data = err.Error()
				handler.ServeHTTP(w, r)
				return
			}
			// Follow the normal flow
			handler.status = http.StatusCreated //TODO: change to OK
			handler.data = true
			w.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(w, r)
			return

		})
	}
}

func SearchUsers(handler *Handler) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			v := mux.Vars(r)
			uid := context.Get(r, "uid")
			users, err := handler.db.SearchUserByName(v["q"], uid.(string))
			if err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Failed to fetch users with query: %v", v["q"]))
				handler.status = http.StatusInternalServerError
				handler.data = err
				handler.ServeHTTP(w, r)
				return
			}

			// Follow the normal flow
			handler.status = http.StatusOK
			handler.data = users
			w.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(w, r)
			return

		})
	}
}

func GetUserContacts(handler *Handler) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u := context.Get(r, "uid")
			users, err := handler.db.FetchUserContacts(u.(string))
			if err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Failed to fetch users with query: %v", u.(string)))
				handler.status = http.StatusInternalServerError
				handler.data = err
				handler.ServeHTTP(w, r)
				return
			}

			// Follow the normal flow
			handler.status = http.StatusOK
			handler.data = users
			w.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(w, r)
			return

		})
	}
}

func GetUser(handler *Handler) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			v := mux.Vars(r)
			uid := v["user"]
			user, err := handler.db.GetUserByUID(uid)
			if err != nil {
				msg := fmt.Sprintf("Failed to fetch users with query: %v ", uid)
				handler.logger.Error(
					fmt.Sprintf("[handler ] msg: %v, err: %v", msg, err))
				handler.status = http.StatusInternalServerError
				handler.data = msg
				handler.ServeHTTP(w, r)
				return
			}

			// Follow the normal flow
			handler.status = http.StatusOK
			handler.data = user
			w.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(w, r)
			return

		})
	}
}

func GetUserBalance(handler *Handler) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			uid := context.Get(r, "uid")
			//TODO: find a more efficient method.. maybe add the eth address in JWT?
			user, err := handler.db.GetUserByUID(uid.(string))
			if err != nil {
				msg := fmt.Sprintf("Failed to fetch users with query: %v", uid)
				handler.logger.Error(
					fmt.Sprintf("[handler ] %v", err))
				handler.status = http.StatusInternalServerError
				handler.data = msg
				handler.ServeHTTP(w, r)
				return
			}
			bal, err := handler.ec.NotaryBalance(user.EthAddress)
			if err != nil {
				msg := fmt.Sprintf("Failed to get notary balance: %v", uid)
				handler.logger.Error(
					fmt.Sprintf("[handler ] %v", msg))
				handler.status = http.StatusInternalServerError
				handler.data = msg
				handler.ServeHTTP(w, r)
				return
			}

			// Follow the normal flow
			handler.status = http.StatusOK
			base18 := math.Pow(10, 18)
			result := big.NewInt(int64(base18))
			handler.data = result.Div(bal, result).Uint64()
			w.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(w, r)
			return

		})
	}
}

func AddContact(handler *Handler) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			v := mux.Vars(r)
			u := context.Get(r, "uid")

			c := UserContact{PUid: u.(string), SUid: v["u"]}
			if err := handler.db.Insert(&c, UserContacts); err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] User insertion to db error! user: %v, err: %v", u, err))
				handler.status = http.StatusInternalServerError
				handler.data = err
				handler.ServeHTTP(w, r)
				return
			}

			// Follow the normal flow
			handler.status = http.StatusOK
			handler.data = true
			w.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(w, r)
			return

		})
	}
}

func CreateCarContract(handler *Handler, contracts chan<- interface{}) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c := &eth.CarContract{}
			if err := decode(r, c); err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Car contract couldn't be parsed! user: %v, err: %v", c, err))
				handler.status = http.StatusBadRequest
				handler.data = "Car contract couldn't be parsed!"
				handler.ServeHTTP(w, r)
				return
			}

			c.CID = int64(time.Now().Unix())
			c.Approved = false // just in case
			//TODO: check if the contract is valid
			if err := handler.db.Insert(c, CarContractCollection); err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Car contract insertion to db error! user: %v, err: %v", c, err))
				handler.status = http.StatusInternalServerError
				handler.data = err
				handler.ServeHTTP(w, r)
				return
			}

			handler.logger.Info(fmt.Sprint("[handler ] Contract successfully saved to db!", (*c).CID))

			uid := context.Get(r, "uid")
			cn := eth.ContractNotification{Type: "new"}
			if cn.NotifyParty = c.Seller; c.Seller == uid {
				cn.NotifyParty = c.Buyer
			}
			uc, err := handler.db.FetchContractByCID(c.CID, uid.(string))
			if err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Can't fetch user for contract! cid: %v, err: %v", c.CID, err))
				handler.status = http.StatusInternalServerError
				handler.data = err
				handler.ServeHTTP(w, r)
				return
			}

			cn.Contract = *uc
			handler.status = http.StatusCreated
			handler.data = c.CID
			contracts <- cn

			h.ServeHTTP(w, r)
			return

		})
	}
}

func SubmitCarContract(handler *Handler, contracts chan<- interface{}) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			v := mux.Vars(r)
			cid := v["cid"]
			//TODO: handle
			i, err := strconv.ParseInt(cid, 10, 64)
			c := handler.db.GetContractByCID(i)
			buyer, seller, err := handler.db.GetContractParticipants(c.Buyer, c.Seller)
			if err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Can't get contract participants db error! cid: %v,", cid))
				handler.status = http.StatusInternalServerError
				handler.data = err
				handler.ServeHTTP(w, r)
				return
			}
			tx, err := handler.ec.CarDeal(c.Hash(), buyer, seller, i)
			if err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Can't write contract to the blockchain! cid: %v, err: %v", cid, err))
				handler.status = http.StatusInternalServerError
				handler.data = "Can't write contract to the blockchain!"
				handler.ServeHTTP(w, r)
				return
			}

			//update contract with hash and approved
			c.ContentHash = c.Hash()
			c.Approved = true
			if err := handler.db.UpdateContract(c); err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Failed to update car contract! contract: %v, err: %v", c, err))
				handler.status = http.StatusInternalServerError
				handler.data = err
				handler.ServeHTTP(w, r)
				return
			}

			uid := context.Get(r, "uid")
			cn := eth.ContractNotification{Type: "approved"}
			if cn.NotifyParty = c.Seller; c.Seller == uid {
				cn.NotifyParty = c.Buyer
			}
			uc, err := handler.db.FetchContractByCID(c.CID, uid.(string))
			if err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Can't fetch user for contract! cid: %v, err: %v", c.CID, err))
				handler.status = http.StatusInternalServerError
				handler.data = err
				handler.ServeHTTP(w, r)
				return
			}

			cn.Contract = *uc
			contracts <- cn
			handler.status = http.StatusOK
			handler.data = tx
			h.ServeHTTP(w, r)

		})
	}
}

func UpdateCarContract(handler *Handler) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c := &eth.CarContract{}
			if err := decode(r, c); err != nil {
				// Set error data and jump to the last handler
				// implemented by *Handler
				handler.status = http.StatusBadRequest
				handler.data = err
				handler.ServeHTTP(w, r)
				return
			}

			if err := handler.db.UpdateContract(c); err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Failed to update car contract! contract: %v, err: %v", c, err))
				handler.status = http.StatusInternalServerError
				handler.data = err
				handler.ServeHTTP(w, r)
				return
			}

			// Follow the normal flow
			handler.status = http.StatusOK
			handler.data = true
			w.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(w, r)
			return

		})
	}
}

func GetUserContracts(handler *Handler) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			v := mux.Vars(r)
			q := v["criteria"]
			u := context.Get(r, "uid")
			users, err := handler.db.FetchUserContracts(u.(string), q)
			if err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Failed to fetch user contracts with query: %v", u.(string)))
				handler.status = http.StatusInternalServerError
				handler.data = err
				handler.ServeHTTP(w, r)
				return
			}

			// Follow the normal flow
			handler.status = http.StatusOK
			handler.data = users
			w.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(w, r)
			return

		})
	}
}

func GetContract(handler *Handler) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			v := mux.Vars(r)
			cid, err := strconv.ParseInt(v["cid"], 10, 64)
			if err != nil {
				msg := fmt.Sprintf("Failed to get cid: %v", cid)
				handler.logger.Error(
					fmt.Sprintf("[handler ] %v", msg))
				handler.status = http.StatusBadRequest
				handler.data = msg
				handler.ServeHTTP(w, r)
				return
			}

			c := handler.db.GetContractByCID(cid)
			if c == nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Failed to fetch user contracts with query: %v", cid))
				handler.status = http.StatusInternalServerError
				handler.data = err
				handler.ServeHTTP(w, r)
				return
			}

			// Follow the normal flow
			handler.status = http.StatusOK
			handler.data = c
			w.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(w, r)
			return

		})
	}
}

func GetContractFieldsList(handler *Handler) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			f := eth.GetContractFields()

			// Follow the normal flow
			handler.status = http.StatusOK
			handler.data = f
			w.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(w, r)
			return

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
				return
			}

			// check if user is valid
			if err := u.OK(); err != nil {
				handler.status = http.StatusInternalServerError
				handler.data = u.OK()
				handler.ServeHTTP(w, r)
				return
			}

			if user, err = handler.db.LoginUserValidation(u); err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Failed to verify user! user: %v, err: %v", u, err))
				handler.status = http.StatusForbidden
				handler.data = err
				handler.ServeHTTP(w, r)
				return
			}

			if user == nil {
				handler.logger.Info(
					fmt.Sprintf("[handler ] Invalid credentials! user: %v", user))
				handler.status = http.StatusForbidden
				handler.data = "Invalid credentials"
				handler.ServeHTTP(w, r)
				return
			}

			//TODO: trigger eth network check

			// Create a map to store user claims
			claims := jwt.StandardClaims{
				IssuedAt: time.Now().Unix(),
				// expires in an hour
				ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
				//TODO: change in production - should be configurable
				Issuer: conf.GetServerAddress(),

				Id: (*user).UID,
			}

			// Create token with claims
			token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

			pvtKey, err := conf.GetPvtKey()
			if err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Unable to get private key! user: %v, err: %v", user, err))
				handler.status = http.StatusInternalServerError
				handler.data = "Server Error" // TODO: may be some meaningful message
				handler.ServeHTTP(w, r)
				return
			}

			tokenString, err := token.SignedString(pvtKey)
			if err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Error while signing the token! user: %v, err: %v", user, err))
				handler.status = http.StatusInternalServerError
				handler.data = "Error while signing the token"
				handler.ServeHTTP(w, r)
				return
			}

			json, err := json.Marshal(&Token{tokenString})
			if err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Error in json marshaling of token string! user: %v, err: %v", user, err))
				handler.status = http.StatusInternalServerError
				handler.data = "Internal Error"
				handler.ServeHTTP(w, r)
				return
			}

			// Follow the normal flow
			handler.status = http.StatusOK
			handler.data = string(json)
			h.ServeHTTP(w, r)

		})
	}
}

func UploadAvatar(handler *Handler, conf *config.Config) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "POST" {
				r.ParseMultipartForm(32 << 20)
				file, fheader, err := r.FormFile("uploadfile")
				if err != nil {
					handler.logger.Error(
						fmt.Sprintf("[handler ] Error while parsing file: %v", err))
					handler.status = http.StatusInternalServerError
					handler.data = "Error while parsing file"
					handler.ServeHTTP(w, r)
					return
				}
				defer file.Close()
				fmt.Fprintf(w, "%v", fheader.Header)
				ext := filepath.Ext(fheader.Filename)
				uid := context.Get(r, "uid").(string)
				filename := conf.GetAvatarDir() + uid + ext
				f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					handler.logger.Error(
						fmt.Sprintf("[handler ] Error while trying to save file: %v", err))
					handler.status = http.StatusInternalServerError
					handler.data = "Error while trying to save file"
					handler.ServeHTTP(w, r)
					return
				}
				defer f.Close()
				_, err = io.Copy(f, file)
				if err != nil {
					handler.logger.Error(
						fmt.Sprintf("[handler ] Error while writing avatar to file: %v", err))
					handler.status = http.StatusInternalServerError
					handler.data = "Error while trying to store file"
					handler.ServeHTTP(w, r)
					return
				}
				u := User{UID: uid, Avatar: uid + ext}
				if err := handler.db.UpdateUser(&u); err != nil {
					handler.logger.Error(
						fmt.Sprintf("[handler ] Failed to update user record!user: %v, err: %v", u, err))
					handler.status = http.StatusInternalServerError
					handler.data = err
					handler.ServeHTTP(w, r)
					return
				}

			} else {
				handler.status = http.StatusMethodNotAllowed
				handler.data = "Method not allowed!"
				h.ServeHTTP(w, r)
			}

		})
	}
}

func DownloadAvatar(handler *Handler, conf *config.Config) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "GET" {
				v := mux.Vars(r)
				uid := v["u"]
				// TODO: check
				u, err := handler.db.GetUserByUID(uid)
				if err != nil {
					handler.logger.Error(
						fmt.Sprintf("[handler ] Error trying to fetch user: %v", err))
					handler.status = http.StatusInternalServerError
					handler.data = " Error trying to fetch user"
					handler.ServeHTTP(w, r)
					return
				}
				filename := u.Avatar
				if len(filename) < 1 {
					handler.logger.Error(
						fmt.Sprintf("[handler ] No avatar filepath for user: %v", u.UID))
					handler.status = http.StatusNoContent
					handler.ServeHTTP(w, r)
					return
				}
				path := filepath.Join(conf.GetAvatarDir(), filename)
				handler.logger.Info(
					fmt.Sprintf("[handler ] Read from file: %v", path))
				f, err := os.OpenFile(path, os.O_RDONLY, 0666)
				if err != nil {
					handler.logger.Error(
						fmt.Sprintf("[handler ] Error while trying to read from file: %v", err))
					handler.status = http.StatusInternalServerError
					handler.data = "Error while trying to read from file"
					handler.ServeHTTP(w, r)
					return
				}
				defer f.Close()
				_, err = io.Copy(w, f)
				if err != nil {
					handler.logger.Error(
						fmt.Sprintf("[handler ] Error while writing avatar response: %v", err))
					handler.status = http.StatusInternalServerError
					handler.data = "Error while writing response"
					handler.ServeHTTP(w, r)
					return
				}
				handler.status = http.StatusOK
				handler.data = "attachment"

			} else {
				handler.status = http.StatusMethodNotAllowed
				handler.data = "Method not allowed!"
				h.ServeHTTP(w, r)
			}

		})
	}
}

/* --------------- Middlewares ---------------- */

//ValidateTokenMiddleware
func ValidateTokenMiddleware(handler *Handler, conf *config.Config) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var token string

			// Get token from the Authorization header
			// format: Authorization: Bearer
			tokens, ok := r.Header["Authorization"]
			if ok && len(tokens) >= 1 {
				token = tokens[0]
				token = strings.TrimPrefix(token, "Bearer ")
			}

			// check if token is empty
			if token == "" {

				handler.logger.Error(
					fmt.Sprint("[handler ] Token string is empty!"))
				handler.status = http.StatusForbidden
				handler.data = "Token string is empty!"
				handler.ServeHTTP(w, r)
				return
			}

			// Parse token (TODO: Assuming pubkey will always be there)
			key, _ := conf.GetPubKey()
			parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					msg := fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					return nil, msg
				}
				return key, nil
			})

			if err != nil {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Token couldn't be parsed! %v", err))
				handler.status = http.StatusForbidden
				handler.data = err.Error()
				handler.ServeHTTP(w, r)
				return
			}
			if parsedToken.Valid {
				id := parsedToken.Claims.(jwt.MapClaims)["jti"]
				context.Set(r, "uid", id)
				handler.status = http.StatusOK
				h.ServeHTTP(w, r)
			} else {
				handler.logger.Error(
					fmt.Sprintf("[handler ] Invalid Token! %v", err))
				handler.status = http.StatusForbidden
				handler.data = "Not allowed"
				handler.ServeHTTP(w, r)
			}

		})
	}
}
