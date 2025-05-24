package routes

import (
	"access-key-management/pkg/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func GetRoutes(r *mux.Router) *mux.Router {
	r.Use(middlewares.LoggingMiddleware)
	//admin apis
	var adminRoutes = r.NewRoute().Subrouter()
	adminRoutes.Use(middlewares.AdminAuthenticator)
	adminRoutes.HandleFunc("/keys", createNewAccessKeys).Methods(http.MethodPost)
	adminRoutes.HandleFunc("/keys/{keyId}", deleteAccessKeys).Methods(http.MethodDelete)
	adminRoutes.HandleFunc("/keys/{keyId}", updateAccessKeys).Methods(http.MethodPut)
	adminRoutes.HandleFunc("/keys", getAllAccessKeys).Methods(http.MethodGet)

	//user apis
	var userRoutes = r.NewRoute().Subrouter()
	userRoutes.Use(middlewares.UserAuthenticator)
	userRoutes.HandleFunc("/keys/{keyId}", fetchAccessKeys).Methods(http.MethodGet)
	userRoutes.HandleFunc("/keys/{keyId}/disable", disableAccessKeys).Methods(http.MethodPost)

	return r
}
