package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/keys", createNewAccessKeys).Methods(http.MethodPost)
	r.HandleFunc("/keys/:keyId", deleteAccessKeys).Methods(http.MethodDelete)
	r.HandleFunc("/keys/:keyId", updateAccessKeys).Methods(http.MethodPut)
	r.HandleFunc("/keys", getAllAccessKeys).Methods(http.MethodGet)

	r.HandleFunc("/keys/{keyId}", fetchAccessKeys).Methods(http.MethodGet)
	r.HandleFunc("/keys/:keyId/disable", disableAccessKeys).Methods(http.MethodPost)

	return r
}
