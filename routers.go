package notary

import (
	"github.com/gorilla/mux"
)

func (n *Notary) muxServer() (router *mux.Router) {

	handler := &Handler{}
	handler.db = n.db
	handler.logger = n.logger

	router = mux.NewRouter()

	router.Handle("/", Adapt(handler, Index(handler), Logging(handler)))

	router.Handle("/sign-up", Adapt(handler, CreateUser(handler, n.email), Logging(handler)))
	// //TODO: technically should be restricted
	router.Handle("/update-info", Adapt(handler, UpdateUserInfo(handler), Logging(handler)))
	router.HandleFunc("/subscribe", ServeWs)
	// // router.HandleFunc("/verify-address", auth.VerifySecondaryAddress)LoginHandler
	router.Handle("/sign-in", Adapt(handler, LoginHandler(handler, n.conf), Logging(handler)))
	router.Handle("/send-invitation", Adapt(handler, NotImplemented(handler), Logging(handler)))

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
