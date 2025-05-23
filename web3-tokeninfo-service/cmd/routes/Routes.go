package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/tokens/{tokenId}", fetchTokenInfo).Methods(http.MethodGet)
	return r
}
