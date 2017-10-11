package notary

import (
	"github.com/NTRYPlatform/ntry-backend/ws"
	"github.com/gorilla/mux"
)

func (n *Notary) muxServer() (router *mux.Router) {

	handler := &Handler{}
	handler.ec = n.ethClient
	handler.db = n.db
	handler.logger = n.logger

	router = mux.NewRouter()

	router.Handle("/", Adapt(handler, Index(handler), Logging(handler))).Methods("GET")
	router.Handle("/sign-up", Adapt(handler, CreateUser(handler, n.email, n.conf), Logging(handler))).Methods("POST")
	router.Handle("/sign-in", Adapt(handler, LoginHandler(handler, n.conf), Logging(handler))).Methods("POST")
	router.Handle("/send-invitation", Adapt(handler, NotImplemented(handler), Logging(handler))).Methods("GET")
	router.Handle("/forgot-password", Adapt(handler, ForgotPassword(handler, n.email, n.conf), Logging(handler))).Methods("POST")
	router.Handle("/change-password", Adapt(handler, ChangePassword(handler), Logging(handler))).Methods("POST")

	//TODO: specify .Headers()
	router.Handle("/user/update", Adapt(handler, UpdateUserInfo(handler), ValidateTokenMiddleware(handler, n.conf), Logging(handler))).Methods("POST")
	router.Handle("/user/upload-avatar", Adapt(handler, UploadAvatar(handler, n.conf), ValidateTokenMiddleware(handler, n.conf), Logging(handler))).Methods("POST")
	//Not authenticated - for client
	router.Handle("/user/get-avatar", Adapt(handler, DownloadAvatar(handler, n.conf), Logging(handler))).
		Queries("u", "{u}").Methods("GET")
	router.Handle("/user/search", Adapt(handler, SearchUsers(handler), ValidateTokenMiddleware(handler, n.conf), Logging(handler))).
		Queries("q", "{q}").Methods("GET")
	router.Handle("/user/get/{user}", Adapt(handler, GetUser(handler), ValidateTokenMiddleware(handler, n.conf), Logging(handler))).Methods("GET")
	router.Handle("/user/balance", Adapt(handler, GetUserBalance(handler), ValidateTokenMiddleware(handler, n.conf), Logging(handler))).Methods("GET")

	router.Handle("/contacts/add", Adapt(handler, AddContact(handler), ValidateTokenMiddleware(handler, n.conf), Logging(handler))).
		Queries("u", "{u}").Methods("GET")
	router.Handle("/contacts/get", Adapt(handler, GetUserContacts(handler), ValidateTokenMiddleware(handler, n.conf), Logging(handler))).Methods("GET")

	router.Handle("/contracts/submit/{cid}", Adapt(handler, SubmitCarContract(handler, n.contracts), ValidateTokenMiddleware(handler, n.conf), Logging(handler))).Methods("GET")
	router.Handle("/contracts/get", Adapt(handler, GetUserContracts(handler), ValidateTokenMiddleware(handler, n.conf), Logging(handler))).Methods("GET")
	router.Handle("/contracts/get/{criteria}", Adapt(handler, GetUserContracts(handler), ValidateTokenMiddleware(handler, n.conf), Logging(handler))).Methods("GET")
	router.Handle("/contracts/fields", Adapt(handler, GetContractFieldsList(handler), ValidateTokenMiddleware(handler, n.conf), Logging(handler))).Methods("GET")
	router.Handle("/contracts/add", Adapt(handler, CreateCarContract(handler, n.contracts), ValidateTokenMiddleware(handler, n.conf), Logging(handler))).Methods("POST")
	router.Handle("/contracts/update", Adapt(handler, UpdateCarContract(handler), ValidateTokenMiddleware(handler, n.conf), Logging(handler))).Methods("POST")
	router.Handle("/contracts/get/{cid}", Adapt(handler, GetContract(handler), ValidateTokenMiddleware(handler, n.conf), Logging(handler))).Methods("GET")

	router.HandleFunc("/subscribe/register/{uid}", ws.ServeRegWs).Methods("GET")
	router.HandleFunc("/subscribe/contract/{uid}", ws.ServeContractWs).Methods("GET")
	router.HandleFunc("/subscribe/contract/approved/{uid}", ws.ServeApprovedContractWs).Methods("GET")

	return router
}
