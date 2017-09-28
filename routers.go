package notary

import (
	"github.com/NTRYPlatform/ntry-backend/ws"
	"github.com/gorilla/mux"
)

func (n *Notary) muxServer() (router *mux.Router) {

	handler := &Handler{}
	handler.db = n.db
	handler.logger = n.logger

	router = mux.NewRouter()

	router.Handle("/", Adapt(handler, Index(handler), Logging(handler))).Methods("GET")
	router.Handle("/sign-up", Adapt(handler, CreateUser(handler, n.email), Logging(handler))).Methods("POST")
	//TODO: specify .Headers()
	// //TODO: technically should be restricted
	router.Handle("/user/update", Adapt(handler, UpdateUserInfo(handler), ValidateTokenMiddleware(handler, n.conf), Logging(handler))).Methods("POST")
	// router.Handle("/update-info", Adapt(handler, UpdateUserInfo(handler), Logging(handler))).Methods("POST")
	router.Handle("/user/search", Adapt(handler, SearchUsers(handler), ValidateTokenMiddleware(handler, n.conf), Logging(handler))).
		Queries("q", "{q}").Methods("GET")

	router.Handle("/contacts/add", Adapt(handler, AddContact(handler), ValidateTokenMiddleware(handler, n.conf), Logging(handler))).
		Queries("u", "{u}").Methods("GET")
	router.Handle("/contacts/get", Adapt(handler, GetUserContacts(handler), ValidateTokenMiddleware(handler, n.conf), Logging(handler))).Methods("GET")

	router.Handle("/contracts/get", Adapt(handler, GetUserContracts(handler), ValidateTokenMiddleware(handler, n.conf), Logging(handler))).Methods("GET")
	router.Handle("/contracts/add", Adapt(handler, CreateCarContract(handler), ValidateTokenMiddleware(handler, n.conf), Logging(handler))).Methods("POST")
	router.Handle("/contracts/update", Adapt(handler, UpdateCarContract(handler), ValidateTokenMiddleware(handler, n.conf), Logging(handler))).Methods("POST")

	router.HandleFunc("/subscribe/register/{uid}", ws.ServeWs).Methods("GET")
	router.Handle("/sign-in", Adapt(handler, LoginHandler(handler, n.conf), Logging(handler))).Methods("POST")
	router.Handle("/send-invitation", Adapt(handler, NotImplemented(handler), Logging(handler))).Methods("GET")

	return router
}

// type Route struct {
// 	Name        string
// 	Method      string
// 	Pattern     string
// 	HandlerFunc http.HandlerFunc
// }

// type Routes []Route

// //TODO: Add others to isolate the routing
// func NewRouter() *mux.Router {

// 	router := mux.NewRouter().StrictSlash(true)
// 	for _, route := range routes {
// 		router.
// 			Methods(route.Method).
// 			Path(route.Pattern).
// 			Name(route.Name).
// 			Handler(route.HandlerFunc)
// 	}

// 	return router
// }

// // define all api routes here
// var routes = Routes{
// 	Route{
// 		"Index",
// 		"GET",
// 		"/",
// 		Index,
// 	},
// }
